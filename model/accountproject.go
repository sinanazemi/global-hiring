package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountProject struct {
  Id     int  `json:"id"`
  Name string `json:"name"`
  Occupation Occupation `json:"occupation"`
  Month Month `json:"month"`
  Year int `json:"year"`
  Url string `json:"url"`
  Description string `json:"description"`
}

func parseAccountProject(data interface{}) (AccountProject, error) {
  result := AccountProject{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountProject', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Occupation, _ = parseOccupation(dataMap["occupation"])
  result.Month, _ = parseMonth(dataMap["month"])
  result.Year = util.ParseInteger(dataMap, "year")
  result.Url = util.ParseString(dataMap, "url")
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountProjects(data interface{}) []AccountProject {
  result := make([]AccountProject, 0)

  tArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'AccountProject's, but not found.\n")
    return result
  }

  for _ , t := range tArr {
    project, err := parseAccountProject(t)
    if (err == nil) {
      result = append(result, project)
    }
  }
  return result
}

func (project *AccountProject) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountProject", "AccountID", project.Id)
}

func (project *AccountProject) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(project.Name)) <= 0 {
    errStr = errStr + "AccountProject.Name is Empty\n"
  }

  if (project.Occupation.Id <= 0) {
    errStr = errStr + "Invalid Occupation in AccountProject\n"
  }

  if (project.Year <= 0 || project.Month.isEmpty()) {
    errStr = errStr + "Invalid date in AccountProject\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (project *AccountProject) insertValidation(session *util.Session) error {
  return project.dataValidation(session)
}

func (project *AccountProject) updateValidation(session *util.Session) error {
  err := project.accountValidation(session)
  if err != nil{
    return err
  }
  return project.dataValidation(session)
}

func (project *AccountProject) deleteValidation(session *util.Session) error {
  return project.accountValidation(session)
}

func loadAccountProjects(session *util.Session) ([]AccountProject, error) {

  query :=
    "SELECT ID, Name, OccupationID, Month, Year, Url, Description " +
    "FROM AccountProject " +
    " WHERE AccountID = $1";

    var result = make([]AccountProject, 0)

    projects, err := util.Select(readAccountProject, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyproject := range projects {

      project, _ := dummyproject.(AccountProject)
      project.Occupation = loadOccupation(project.Occupation.Id)
      project.Month = loadMonth(project.Month.Value)
      result = append(result, project)
    }

    return result, nil
}

func readAccountProject(rows *sql.Rows) (interface{}, error) {

    var project AccountProject = AccountProject{}
    err := rows.Scan(&project.Id, &project.Name, &project.Occupation.Id, &project.Month.Value, &project.Year, &project.Url, &project.Description)

    return project, err
}

func (project *AccountProject) save(session *util.Session) error {
  if project.Id <= 0 {
    return project.saveNew(session)
  }
  return project.saveUpdate(session)
}

func (project *AccountProject) saveNew(session *util.Session) error {

  err := project.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountProject " +
    "(Title, OccupationID, Month, Year, Url, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6) " +
    "returning ID"

  id, err := util.Insert(query, project.Name, project.Occupation.Id, project.Month.Value, project.Year, project.Url, project.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  project.Id = id
  return nil
}

func (project *AccountProject) saveUpdate(session *util.Session) error {
  err := project.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountProject " +
    "SET " +
    "Name = $1, " +
    "OccupationID = $2, " +
    "Month = $3, " +
    "Year = $4, " +
    "Url = $5, " +
    "Description = $6 " +
    "WHERE ID = $7 "

  err = util.Update(query, project.Name, project.Occupation.Id, project.Month.Value, project.Year, project.Url, project.Description, project.Id)

  return err

}

func (project *AccountProject) delete(session *util.Session) error {
  err := project.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountProject " +
    "WHERE ID = $1 "

  err = util.Update(query, project.Id)

  return err
}

func getProjectStrength(projects []AccountProject) int {

  //adding project +10

  if (projects == nil) {
    return 0;
  }
  size := len(projects)
  if(size < 1) {
    return 0
  }
  return 10
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveProject(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  projectMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON project", http.StatusBadRequest}
  }

  project, _ := parseAccountProject(projectMap)

  err = project.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving project", http.StatusBadRequest}
  }

  return project, nil
}

func DeleteProject(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  projectMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON project", http.StatusBadRequest}
  }

  project, _ := parseAccountProject(projectMap)

  err = project.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting project", http.StatusBadRequest}
  }

  return project, nil
}

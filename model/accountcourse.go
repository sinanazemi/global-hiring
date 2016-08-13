package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountCourse struct {
  Id     int  `json:"id"`
  Name string `json:"name"`
  Number string `json:"number"`
  Occupation Occupation `json:"occupation"`
  Description string `json:"description"`
}

func parseAccountCourse(data interface{}) (AccountCourse, error) {
  result := AccountCourse{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountCourse', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Number = util.ParseString(dataMap, "number")
  result.Occupation, _ = parseOccupation(dataMap["occupation"])
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountCourses(data interface{}) []AccountCourse {
  result := make([]AccountCourse, 0)

  cArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'AccountCourse's, but not found.\n")
    return result
  }

  for _ , c := range cArr {
    course, err := parseAccountCourse(c)
    if (err == nil) {
      result = append(result, course)
    }
  }
  return result
}

func (course *AccountCourse) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountCourse", "AccountID", course.Id)
}

func (course *AccountCourse) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(course.Name)) <= 0 {
    errStr = errStr + "AccountCourse.Name is Empty\n"
  }

  if len(strings.TrimSpace(course.Number)) <= 0 {
    errStr = errStr + "AccountCourse.Number is Empty\n"
  }

  if (course.Occupation.Id <= 0) {
    errStr = errStr + "Invalid Occupation in AccountCourse\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (course *AccountCourse) insertValidation(session *util.Session) error {
  return course.dataValidation(session)
}

func (course *AccountCourse) updateValidation(session *util.Session) error {
  err := course.accountValidation(session)
  if err != nil{
    return err
  }
  return course.dataValidation(session)
}

func (course *AccountCourse) deleteValidation(session *util.Session) error {
  return course.accountValidation(session)
}

func loadAccountCourses(session *util.Session) ([]AccountCourse, error) {

  query :=
    "SELECT ID, Name, Number, OccupationID, Description " +
    "FROM AccountCourse " +
    " WHERE AccountID = $1";

    var result = make([]AccountCourse, 0)

    courses, err := util.Select(readAccountCourse, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyCourse := range courses {

      course, _ := dummyCourse.(AccountCourse)
      course.Occupation = loadOccupation(course.Occupation.Id)
      result = append(result, course)
    }

    return result, nil
}

func readAccountCourse(rows *sql.Rows) (interface{}, error) {

    var course AccountCourse = AccountCourse{}
    err := rows.Scan(&course.Id, &course.Name, &course.Number, &course.Occupation.Id, &course.Description)

    return course, err
}

func (course *AccountCourse) save(session *util.Session) error {
  if course.Id <= 0 {
    return course.saveNew(session)
  }
  return course.saveUpdate(session)
}

func (course *AccountCourse) saveNew(session *util.Session) error {

  err := course.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountCourse " +
    "(Name, Number, OccupationID, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5) " +
    "returning ID"

  id, err := util.Insert(query, course.Name, course.Number, course.Occupation.Id, course.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  course.Id = id
  return nil
}

func (course *AccountCourse) saveUpdate(session *util.Session) error {
  err := course.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountCourse " +
    "SET " +
    "Name = $1, " +
    "Number = $2, " +
    "OccupationID = $3, " +
    "Description = $4 " +
    "WHERE ID = $5 "

  err = util.Update(query, course.Name, course.Number, course.Occupation.Id, course.Description, course.Id)

  return err

}

func (course *AccountCourse) delete(session *util.Session) error {
  err := course.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountCourse " +
    "WHERE ID = $1 "

  err = util.Update(query, course.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveCourse(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  courseMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Course", http.StatusBadRequest}
  }

  course, _ := parseAccountCourse(courseMap)

  err = course.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Course", http.StatusBadRequest}
  }

  return course, nil
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  courseMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Course", http.StatusBadRequest}
  }

  course, _ := parseAccountCourse(courseMap)

  err = course.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Course", http.StatusBadRequest}
  }

  return course, nil
}

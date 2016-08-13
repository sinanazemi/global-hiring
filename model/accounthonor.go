package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountHonor struct {
  Id     int  `json:"id"`
  Title string `json:"title"`
  Occupation Occupation `json:"occupation"`
  Month Month `json:"month"`
  Year int `json:"year"`
  Description string `json:"description"`
}

func parseAccountHonor(data interface{}) (AccountHonor, error) {
  result := AccountHonor{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountHonor', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Title = util.ParseString(dataMap, "title")
  result.Occupation, _ = parseOccupation(dataMap["occupation"])
  result.Month, _ = parseMonth(dataMap["month"])
  result.Year = util.ParseInteger(dataMap, "year")
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountHonors(data interface{}) []AccountHonor {
  result := make([]AccountHonor, 0)

  hArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'AccountHonor's, but not found.\n")
    return result
  }

  for _ , h := range hArr {
    honor, err := parseAccountHonor(h)
    if (err == nil) {
      result = append(result, honor)
    }
  }
  return result
}

func (honor *AccountHonor) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountHonor", "AccountID", honor.Id)
}

func (honor *AccountHonor) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(honor.Title)) <= 0 {
    errStr = errStr + "AccountHonor.Title is Empty\n"
  }

  if (honor.Occupation.Id <= 0) {
    errStr = errStr + "Invalid Occupation in AccountHonor\n"
  }

  if (honor.Year <= 0 || honor.Month.isEmpty()) {
    errStr = errStr + "Invalid date in AccountHonor\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (honor *AccountHonor) insertValidation(session *util.Session) error {
  return honor.dataValidation(session)
}

func (honor *AccountHonor) updateValidation(session *util.Session) error {
  err := honor.accountValidation(session)
  if err != nil{
    return err
  }
  return honor.dataValidation(session)
}

func (honor *AccountHonor) deleteValidation(session *util.Session) error {
  return honor.accountValidation(session)
}

func loadAccountHonors(session *util.Session) ([]AccountHonor, error) {

  query :=
    "SELECT ID, Title,  OccupationID, Month, Year, Description " +
    "FROM AccountHonor " +
    " WHERE AccountID = $1";

    var result = make([]AccountHonor, 0)

    honors, err := util.Select(readAccountHonor, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyhonor := range honors {

      honor, _ := dummyhonor.(AccountHonor)
      honor.Occupation = loadOccupation(honor.Occupation.Id)
      honor.Month = loadMonth(honor.Month.Value)
      result = append(result, honor)
    }

    return result, nil
}

func readAccountHonor(rows *sql.Rows) (interface{}, error) {

    var honor AccountHonor = AccountHonor{}
    err := rows.Scan(&honor.Id, &honor.Title, &honor.Occupation.Id, &honor.Month.Value, &honor.Year, &honor.Description)

    return honor, err
}

func (honor *AccountHonor) save(session *util.Session) error {
  if honor.Id <= 0 {
    return honor.saveNew(session)
  }
  return honor.saveUpdate(session)
}

func (honor *AccountHonor) saveNew(session *util.Session) error {

  err := honor.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountHonor " +
    "(Title, OccupationID, Month, Year, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6) " +
    "returning ID"

  id, err := util.Insert(query, honor.Title, honor.Occupation.Id, honor.Month.Value, honor.Year, honor.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  honor.Id = id
  return nil
}

func (honor *AccountHonor) saveUpdate(session *util.Session) error {
  err := honor.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountHonor " +
    "SET " +
    "Title = $1, " +
    "OccupationID = $2, " +
    "Month = $3, " +
    "Year = $4, " +
    "Description = $5 " +
    "WHERE ID = $6 "

  err = util.Update(query, honor.Title, honor.Occupation.Id, honor.Month.Value, honor.Year, honor.Description, honor.Id)

  return err

}

func (honor *AccountHonor) delete(session *util.Session) error {
  err := honor.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountHonor " +
    "WHERE ID = $1 "

  err = util.Update(query, honor.Id)

  return err
}

func getHonorStrength(honors []AccountHonor) int {

  //adding rewards +10

  if (honors == nil) {
    return 0;
  }
  size := len(honors)
  if(size < 1) {
    return 0
  }
  return 10
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveHonor(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  honorMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON honor", http.StatusBadRequest}
  }

  honor, _ := parseAccountHonor(honorMap)

  err = honor.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving honor", http.StatusBadRequest}
  }

  return honor, nil
}

func DeleteHonor(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  honorMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON honor", http.StatusBadRequest}
  }

  honor, _ := parseAccountHonor(honorMap)

  err = honor.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting honor", http.StatusBadRequest}
  }

  return honor, nil
}

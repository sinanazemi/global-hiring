package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountWork struct {
  Id     int  `json:"id"`
  Company string `json:"company"`
  Location string `json:"location"`
  Title string `json:"title"`
  Role WorkRole `json:"role"`
  FromMonth Month `json:"frommonth"`
  FromYear int `json:"fromyear"`
  ToMonth Month `json:"tomonth"`
  ToYear int `json:"toyear"`
  Currently bool `json:"currently"`
  Description string `json:"description"`
}

func parseAccountWork(data interface{}) (AccountWork, error) {
  result := AccountWork{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountWork', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Company = util.ParseString(dataMap, "company")
  result.Location = util.ParseString(dataMap, "location")
  result.Title = util.ParseString(dataMap, "title")

  result.Role, _ = parseWorkRole(dataMap["role"])

  result.FromMonth, _ = parseMonth(dataMap["frommonth"])
  result.FromYear = util.ParseInteger(dataMap, "fromyear")

  result.ToMonth, _ = parseMonth(dataMap["tomonth"])
  result.ToYear = util.ParseInteger(dataMap, "toyear")

  result.Currently = util.ParseBool(dataMap, "currently")
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountWorks(data interface{}) []AccountWork {
  result := make([]AccountWork, 0)

  wArr, ok := data.([]interface{})
  if (!ok) {
    print("lokking for a 'm[]interface{}' to parse an array of 'AccountWork's, but not found.\n")
    return result
  }

  for _ , w := range wArr {
    work, err := parseAccountWork(w)
    if (err == nil) {
      result = append(result, work)
    }
  }
  return result
}

func (work *AccountWork) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountWork", "AccountID", work.Id)
}

func (work *AccountWork) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(work.Company)) <= 0 {
    errStr = errStr + "AccountWork.Company is Empty\n"
  }

  if len(strings.TrimSpace(work.Location)) <= 0 {
    errStr = errStr + "AccountWork.Location is Empty\n"
  }

  if len(strings.TrimSpace(work.Title)) <= 0 {
    errStr = errStr + "AccountWork.Title is Empty\n"
  }

  if (work.Role.isEmpty()) {
    errStr = errStr + "Invalid Role in AccountWork\n"
  }

  if (work.Currently) {
    work.ToYear = 0
    work.ToMonth = getEmptyMonth()
  }

  if (work.FromYear <= 0 || work.FromMonth.isEmpty()) {
    errStr = errStr + "Invalid start date in AccountWork\n"
  }

  var from = (work.FromYear * 12) + work.FromMonth.Value
  var to = (work.ToYear * 12) + work.ToMonth.Value

  if (!work.Currently && to < from) {
    errStr = errStr + "Invalid relation between dates in AccountWork\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (work *AccountWork) insertValidation(session *util.Session) error {
  return work.dataValidation(session)
}

func (work *AccountWork) updateValidation(session *util.Session) error {
  err := work.accountValidation(session)
  if err != nil{
    return err
  }
  return work.dataValidation(session)
}

func (work *AccountWork) deleteValidation(session *util.Session) error {
  return work.accountValidation(session)
}

func loadAccountWorks(session *util.Session) ([]AccountWork, error) {

  query :=
    "SELECT ID, Company, Location, Title, Role, FromMonth, FromYear, ToMonth, ToYear, currently, Description " +
    "FROM AccountWork " +
    " WHERE AccountID = $1";

    var result = make([]AccountWork, 0)

    works, err := util.Select(readAccountWork, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyWork := range works {

      work, _ := dummyWork.(AccountWork)

      work.Role = loadWorkRole(work.Role.Value)
      work.FromMonth = loadMonth(work.FromMonth.Value)
      work.ToMonth = loadMonth(work.ToMonth.Value)

      result = append(result, work)
    }

    return result, nil
}

func readAccountWork(rows *sql.Rows) (interface{}, error) {

    var work AccountWork = AccountWork{}
    err := rows.Scan(&work.Id, &work.Company, &work.Location, &work.Title, &work.Role.Value, &work.FromMonth.Value, &work.FromYear, &work.ToMonth.Value, &work.ToYear, &work.Currently, &work.Description)

    return work, err
}

func (work *AccountWork) save(session *util.Session) error {
  if work.Id <= 0 {
    return work.saveNew(session)
  }
  return work.saveUpdate(session)
}

func (work *AccountWork) saveNew(session *util.Session) error {

  err := work.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountWork " +
    "(Company, Location, Title, Role, FromMonth, FromYear, ToMonth, ToYear, currently, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) " +
    "returning ID"

  id, err := util.Insert(query, work.Company, work.Location, work.Title, work.Role.Value, work.FromMonth.Value, work.FromYear, work.ToMonth.Value, work.ToYear, work.Currently, work.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  work.Id = id
  return nil
}

func (work *AccountWork) saveUpdate(session *util.Session) error {
  err := work.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountWork " +
    "SET " +
    "Company = $1, " +
    "Location = $2, " +
    "Title = $3, " +
    "Role = $4, " +
    "FromMonth = $5, " +
    "FromYear = $6, " +
    "ToMonth = $7, " +
    "ToYear = $8, " +
    "currently = $9, " +
    "Description = $10 " +
    "WHERE ID = $11 "

  err = util.Update(query, work.Company, work.Location, work.Title, work.Role.Value, work.FromMonth.Value, work.FromYear, work.ToMonth.Value, work.ToYear, work.Currently, work.Description, work.Id)

  return err

}

func (work *AccountWork) delete(session *util.Session) error {
  err := work.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountWork " +
    "WHERE ID = $1 "

  err = util.Update(query, work.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveWork(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  workMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Work", http.StatusBadRequest}
  }

  work, _ := parseAccountWork(workMap)

  err = work.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Work", http.StatusBadRequest}
  }

  return work, nil
}

func DeleteWork(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  workMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Work", http.StatusBadRequest}
  }

  work, _ := parseAccountWork(workMap)

  err = work.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Work", http.StatusBadRequest}
  }

  return work, nil
}

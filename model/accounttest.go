package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountTest struct {
  Id     int  `json:"id"`
  Name string `json:"name"`
  Occupation Occupation `json:"occupation"`
  Month Month `json:"month"`
  Year int `json:"year"`
  Score float64 `json:"score"`
  Description string `json:"description"`
}

func parseAccountTest(data interface{}) (AccountTest, error) {
  result := AccountTest{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountTest', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Occupation, _ = parseOccupation(dataMap["occupation"])
  result.Month, _ = parseMonth(dataMap["month"])
  result.Year = util.ParseInteger(dataMap, "year")
  result.Score = util.ParseFloat(dataMap, "score")
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountTests(data interface{}) []AccountTest {
  result := make([]AccountTest, 0)

  tArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'AccountTest's, but not found.\n")
    return result
  }

  for _ , t := range tArr {
    test, err := parseAccountTest(t)
    if (err == nil) {
      result = append(result, test)
    }
  }
  return result
}

func (test *AccountTest) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountTest", "AccountID", test.Id)
}

func (test *AccountTest) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(test.Name)) <= 0 {
    errStr = errStr + "AccountTest.Name is Empty\n"
  }

  if (test.Occupation.Id <= 0) {
    errStr = errStr + "Invalid Occupation in AccountTest\n"
  }

  if (test.Year <= 0 || test.Month.isEmpty()) {
    errStr = errStr + "Invalid date in AccountTest\n"
  }

  if (test.Score < 0 ) {
    errStr = errStr + "Invalid Score in AccountTest\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (test *AccountTest) insertValidation(session *util.Session) error {
  return test.dataValidation(session)
}

func (test *AccountTest) updateValidation(session *util.Session) error {
  err := test.accountValidation(session)
  if err != nil{
    return err
  }
  return test.dataValidation(session)
}

func (test *AccountTest) deleteValidation(session *util.Session) error {
  return test.accountValidation(session)
}

func loadAccountTests(session *util.Session) ([]AccountTest, error) {

  query :=
    "SELECT ID, Name, OccupationID, Month, Year, Score, Description " +
    "FROM AccountTest " +
    " WHERE AccountID = $1";

    var result = make([]AccountTest, 0)

    tests, err := util.Select(readAccountTest, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummytest := range tests {

      test, _ := dummytest.(AccountTest)
      test.Occupation = loadOccupation(test.Occupation.Id)
      test.Month = loadMonth(test.Month.Value)
      result = append(result, test)
    }

    return result, nil
}

func readAccountTest(rows *sql.Rows) (interface{}, error) {

    var test AccountTest = AccountTest{}
    err := rows.Scan(&test.Id, &test.Name, &test.Occupation.Id, &test.Month.Value, &test.Year, &test.Score, &test.Description)

    return test, err
}

func (test *AccountTest) save(session *util.Session) error {
  if test.Id <= 0 {
    return test.saveNew(session)
  }
  return test.saveUpdate(session)
}

func (test *AccountTest) saveNew(session *util.Session) error {

  err := test.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountTest " +
    "(Title, OccupationID, Month, Year, Score, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6) " +
    "returning ID"

  id, err := util.Insert(query, test.Name, test.Occupation.Id, test.Month.Value, test.Year, test.Score, test.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  test.Id = id
  return nil
}

func (test *AccountTest) saveUpdate(session *util.Session) error {
  err := test.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountTest " +
    "SET " +
    "Name = $1, " +
    "OccupationID = $2, " +
    "Month = $3, " +
    "Year = $4, " +
    "Score = $5 " +
    "Description = $6 " +
    "WHERE ID = $7 "

  err = util.Update(query, test.Name, test.Occupation.Id, test.Month.Value, test.Year, test.Score, test.Description, test.Id)

  return err

}

func (test *AccountTest) delete(session *util.Session) error {
  err := test.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountTest " +
    "WHERE ID = $1 "

  err = util.Update(query, test.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveTest(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  testMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON test", http.StatusBadRequest}
  }

  test, _ := parseAccountTest(testMap)

  err = test.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving test", http.StatusBadRequest}
  }

  return test, nil
}

func DeleteTest(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  testMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON test", http.StatusBadRequest}
  }

  test, _ := parseAccountTest(testMap)

  err = test.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting test", http.StatusBadRequest}
  }

  return test, nil
}

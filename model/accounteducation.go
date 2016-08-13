package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountEducation struct {
  Id int `json:"id"`
  School string `json:"school"`
  FromDate int `json:"fromdate"`
  ToDate int `json:"todate"`
  Field string `json:"field"`
  Grade float64 `json:"grade"`

  Degree Degree `json:"degree"`
}

func parseAccountEducation(data interface{}) (AccountEducation, error) {
  result := AccountEducation{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountEducation', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.School = util.ParseString(dataMap, "school")
  result.FromDate = util.ParseInteger(dataMap, "fromdate")
  result.ToDate = util.ParseInteger(dataMap, "todate")
  result.Field = util.ParseString(dataMap, "field")
  result.Grade = util.ParseFloat(dataMap, "grade")
  result.Degree, _ = parseDegree(dataMap["degree"])

  return result, nil
}

func parseAccountEducations(data interface{}) []AccountEducation {
  result := make([]AccountEducation, 0)

  eArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a 'm[]interface{}' to parse an array of 'AccountEducation's, but not found.\n")
    return result
  }

  for _ , e := range eArr {
    edu, err := parseAccountEducation(e)
    if (err == nil) {
      result = append(result, edu)
    }
  }
  return result
}

func (ace *AccountEducation) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountEducation", "AccountID", ace.Id)
}

func (ace *AccountEducation) dataValidation(session *util.Session) error {
  errStr := ""

  if len(strings.TrimSpace(ace.School)) <= 0 {
    errStr = errStr + "AccountSkill.School is Empty\n"
  }

  if (ace.FromDate <= 0 || ace.ToDate <= 0 || ace.FromDate > ace.ToDate) {
    errStr = errStr + "Invalid dates in AccountSkill\n"
  }

  if len(strings.TrimSpace(ace.Field)) <= 0 {
    errStr = errStr + "AccountSkill.Field is Empty\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (ace *AccountEducation) insertValidation(session *util.Session) error {
  return ace.dataValidation(session)
}

func (ace *AccountEducation) updateValidation(session *util.Session) error {
  err := ace.accountValidation(session)
  if err != nil{
    return err
  }
  return ace.dataValidation(session)
}

func (ace *AccountEducation) deleteValidation(session *util.Session) error {
  return ace.accountValidation(session)
}

func loadAccountEducations(session *util.Session) ([]AccountEducation, error) {

  query :=
    "SELECT ID, School, FromDate, ToDate, Field, Grade, degreeID " +
    "FROM AccountEducation " +
    " WHERE AccountID = $1";

    var result = make([]AccountEducation, 0)

    educations, err := util.Select(readAccountEducation, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyEdu := range educations {
      education, _ := dummyEdu.(AccountEducation)
      education.Degree = loadDegree(education.Degree.Id)
      result = append(result, education)
    }

    return result, nil
}

func readAccountEducation(rows *sql.Rows) (interface{}, error) {

    var edu AccountEducation = AccountEducation{}
    err := rows.Scan(&edu.Id, &edu.School, &edu.FromDate, &edu.ToDate, &edu.Field, &edu.Grade, &edu.Degree.Id)

    return edu, err
}

func (edu *AccountEducation) save(session *util.Session) error {
  if edu.Id <= 0 {
    return edu.saveNew(session)
  }
  return edu.saveUpdate(session)
}

func (edu *AccountEducation) saveNew(session *util.Session) error {

  err := edu.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountEducation" +
    "(School, FromDate, ToDate, Field, Grade, DegreeID, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7) " +
    "returning AccountEducation.ID"

  id, err := util.Insert(query, edu.School, edu.FromDate, edu.ToDate, edu.Field, edu.Grade, edu.Degree.Id, session.GetAccountID())

  print("inserted id for accountEducation is:")
  print(id)
  print("\n")

  if err != nil {
    return err
  }

  print("edu.Id = ")
  print(id)
  print("\n")
  edu.Id = id
  return nil
}

func (edu *AccountEducation) saveUpdate(session *util.Session) error {
  err := edu.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountEducation " +
    "SET " +
    "School = $1," +
    " FromDate = $2," +
    " ToDate = $3," +
    " Field = $4," +
    " Grade = $5," +
    " DegreeID = $6 " +
    "WHERE ID = $7 "

  err = util.Update(query, edu.School, edu.FromDate, edu.ToDate, edu.Field, edu.Grade, edu.Degree.Id, edu.Id)

  return err

}

func (edu *AccountEducation) delete(session *util.Session) error {
  err := edu.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountEducation " +
    "WHERE ID = $1 "

  err = util.Update(query, edu.Id)

  return err
}

func getEducationStrength(educations []AccountEducation) int {

  //adding first education +15
  //adding second education +10
  //adding third education and +5
  //(more education doesn't change)

  if (educations == nil) {
    return 0;
  }
  size := len(educations)
  if(size == 0) {
    return 0
  }
  if(size == 1) {
    return 15
  }
  if(size == 2) {
    return 25
  }
  return 30
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveEducation(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  eduMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Education", http.StatusBadRequest}
  }

  edu, _ := parseAccountEducation(eduMap)

  err = edu.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Education", http.StatusBadRequest}
  }


  print("sending accountEducation - id is:")
  print(edu.Id)
  print("\n")

  return edu, nil
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  eduMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Education", http.StatusBadRequest}
  }

  edu, _ := parseAccountEducation(eduMap)

  err = edu.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Education", http.StatusBadRequest}
  }

  return edu, nil
}

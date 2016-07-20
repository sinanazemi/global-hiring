package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountEducation struct {
  Id     int  `json:"id"`
  School string `json:"school"`
  FromDate int `json:"fromdate"`
  ToDate int `json:"todate"`
  Field string `json:"field"`
  Grade float64 `json:"grade"`

  Degree Degree `json:"degree"`
}

func parseAccountEducation(dataMap map[string]interface{}) (AccountEducation, error) {
  result := AccountEducation{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.School = util.ParseString(dataMap, "school")
  result.FromDate = util.ParseInteger(dataMap, "fromdate")
  result.ToDate = util.ParseInteger(dataMap, "todate")
  result.Field = util.ParseString(dataMap, "field")

  result.Grade = util.ParseFloat(dataMap, "grade")

  degreeMap := dataMap["degree"].(map[string]interface{})
  degree, _ := parseDegree(degreeMap)
  result.Degree = degree
  return result, nil
}

func parseAccountEducations(edusArr []interface{}) []AccountEducation {
  result := make([]AccountEducation, 0)

  for _ , edu := range edusArr {
    emap := edu.(map[string]interface{})
    education, err := parseAccountEducation(emap)
    if (err == nil) {
      result = append(result, education)
    }
  }
  return result
}

func (ace AccountEducation) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountEducation", "AccountID", ace.Id)
}

func (ace AccountEducation) dataValidation(session *util.Session) error {
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

func (ace AccountEducation) insertValidation(session *util.Session) error {
  return ace.dataValidation(session)
}

func (ace AccountEducation) updateValidation(session *util.Session) error {
  err := ace.accountValidation(session)
  if err != nil{
    return err
  }
  return ace.dataValidation(session)
}

func (ace AccountEducation) deleteValidation(session *util.Session) error {
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

func (ace AccountEducation) save(session *util.Session) error {
  if ace.Id <= 0 {
    return ace.saveNew(session)
  }
  return ace.saveUpdate(session)
}

func (edu AccountEducation) saveNew(session *util.Session) error {

  err := edu.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountEducation" +
    "(School, FromDate, ToDate, Field, Grade, DegreeID, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7) " +
    "returning ID"

  id, err := util.Insert(query, edu.School, edu.FromDate, edu.ToDate, edu.Field, edu.Grade, edu.Degree.Id, session.GetAccountID())

  if err != nil {
    return err
  }

  edu.Id = id
  return nil
}

func (edu AccountEducation) saveUpdate(session *util.Session) error {
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

func (edu AccountEducation) delete(session *util.Session) error {
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

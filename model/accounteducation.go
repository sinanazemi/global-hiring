package model

import (
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

func (edu AccountEducation) save(session *util.Session) error {
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

package model

import (
  "github.com/sinanazemi/global-hiring/util"
)

type Education struct {
  Id     int  `json:"id"`
  School string `json:"school"`
  FromDate int `json:"fromdate"`
  ToDate int `json:"todate"`
  Field string `json:"field"`
  Grade float64 `json:"grade"`

  Degree Degree `json:"degree"`
}

func parseEducation(dataMap map[string]interface{}) (Education, error) {
  result := Education{}

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

func parseEducations(edusArr []interface{}) []Education {
  result := make([]Education, 0)

  for _ , edu := range edusArr {
    emap := edu.(map[string]interface{})
    education, err := parseEducation(emap)
    if (err == nil) {
      result = append(result, education)
    }
  }
  return result
}

func (edu Education) save(account Account) error {
  query :=
    "INSERT INTO AccountEducation" +
    "(School, FromDate, ToDate, Field, Grade, DegreeID, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7) " +
    "returning ID"

  id, err := util.Insert(query, edu.School, edu.FromDate, edu.ToDate, edu.Field, edu.Grade, edu.Degree.Id, account.Id)

  if err != nil {
    return err
  }

  edu.Id = id
  return nil
}

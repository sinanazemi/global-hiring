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

func parseEducation(dataMap map[string]interface{}) (interface{}, error) {
  result := Education{}

  result.School = dataMap["school"].(string)
  result.FromDate = int(dataMap["fromdate"].(float64))
  result.ToDate = int(dataMap["todate"].(float64))
  result.Field = dataMap["field"].(string)
  if (dataMap["fromdate"] != nil) {
    result.Grade = dataMap["grade"].(float64)
  }

  degreeMap := dataMap["degree"].(map[string]interface{})
  degree := parseDegreeReturn(degreeMap)
  result.Degree = degree

  return result, nil
}

func parseEducationReturn(eduMap map[string]interface{}) Education {
  edu, _ := parseEducation(eduMap)
  result := edu.(Education)
  return result
}

func parseEducationsReturn(edusArr []interface{}) []Education {
  result := make([]Education, 0)

  for _ , edu := range edusArr {
    emap := edu.(map[string]interface{})
    education := parseEducationReturn(emap)
    result = append(result, education)
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

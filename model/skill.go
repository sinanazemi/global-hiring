package model

import (
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Skill struct {
  Id int `json:"id"`
  Name string `json:"name"`
  MainServiceID int `json:"mainserviceid"`
  Profeciency SkillProfeciency `json:"profeciency"` // dummy field, just for leva
  IsSelected bool `json:"isselected"` // dummy field, just for leva
}

func parseSkill(data interface{}) (Skill, error) {
  result := Skill{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'Skill', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.MainServiceID = util.ParseInteger(dataMap, "mainserviceid")
  result.Profeciency, _ = parseSkillProfeciency(dataMap["profeciency"])
  result.IsSelected = util.ParseBool(dataMap, "isselected")

  return result, nil
}

func parseSkills(data interface{}) []Skill {
  result := make([]Skill, 0)

  sArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'Skill's, but not found.\n")
    return result
  }

  for _ , s := range sArr {
    skill, err := parseSkill(s)
    if (err == nil) {
      result = append(result, skill)
    }
  }
  return result
}

func getSkills(serviceID int) []Skill {

  // list of all of services
	var result = make([]Skill, 0)

  skills, err := util.Select(readSkill, "select * from Skill where MainServiceID = $1" , serviceID)

  if err != nil {
    return result
  }

  for _, dummySkill := range skills {

    skill, ok := dummySkill.(Skill)
    if !ok {
        // service was not of type MainService. The assertion failed
        return make([]Skill, 0)
    }
    // service is of type MainService
    result = append(result, skill)
  }
  return result

}

func readSkill(rows *sql.Rows) (interface{}, error) {

  var skill Skill = Skill{}
  err := rows.Scan(&skill.Id, &skill.Name, &skill.MainServiceID)

  return skill, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetSkills(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  id := util.GetID(r)
  return getSkills(id), nil
}

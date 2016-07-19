package model

import (
  "database/sql"
  "github.com/sinanazemi/global-hiring/util"
)

type Skill struct {
  Id int `json:"id"`
  Name string `json:"name"`
  MainServiceID int `json:"mainserviceid"`

  IsSelected bool `json:"isselected"`
  Profeciency string
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
  skill.IsSelected = false

  return skill, err
}

func parseSkill(dataMap map[string]interface{}) (Skill, error) {
  result := Skill{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.MainServiceID = util.ParseInteger(dataMap, "mainserviceid")
  result.IsSelected = util.ParseBool(dataMap, "isselected")
  result.Profeciency = util.ParseString(dataMap, "profeciency")

  return result, nil
}

func parseSkills(skillsArr []interface{}) []Skill {
  result := make([]Skill, 0)

  for _ , s := range skillsArr {
    smap := s.(map[string]interface{})
    skill, err := parseSkill(smap)
    if err == nil {
      result = append(result, skill)
    }
  }
  return result
}

func (skill Skill) save(account Account) error {
  if !skill.IsSelected {
    return nil
  }

  accountSkill := AccountSkill{Skill: skill, Account: account, Profeciency : skill.Profeciency}
  return accountSkill.save()
}

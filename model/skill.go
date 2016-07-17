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

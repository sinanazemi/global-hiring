package model

import (
  "strconv"
)

type Skill struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Proficiency int `json:"profieciency"`
}

func getSkills(serviceID int) []Skill {

  print("get skills " + strconv.Itoa(serviceID) + "\n")

	var skills = make([]Skill, 0)

  // bootstrap some data
  skills = append(skills, Skill{1, "Skill 1 for service " + strconv.Itoa(serviceID), 0})
  skills = append(skills, Skill{1, "Skill 2 for service " + strconv.Itoa(serviceID), 0})
  skills = append(skills, Skill{1, "Skill 3 for service " + strconv.Itoa(serviceID), 0})
  skills = append(skills, Skill{1, "Skill 4 for service " + strconv.Itoa(serviceID), 0})
  skills = append(skills, Skill{1, "Skill 5 for service " + strconv.Itoa(serviceID), 0})

  return skills

}

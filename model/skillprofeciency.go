package model

import(
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type SkillProfeciency struct{
  Value string `json:"value"`
  Name string `json:"name"`
}

var skillProfeciencies []SkillProfeciency = make([]SkillProfeciency, 0)

var emptySkillProfeciency SkillProfeciency = SkillProfeciency{"", ""}

func (s SkillProfeciency) isEmpty() bool {
  return len(s.Value) <= 0
}

func (s SkillProfeciency) clone() SkillProfeciency {
  return SkillProfeciency{s.Value, s.Name}
}

func getEmptySkillProfeciency() SkillProfeciency {
  return emptySkillProfeciency.clone()
}

func loadSkillProfeciencies() {
  if len(skillProfeciencies) > 0 {
    return
  }

  skillProfeciencies = append(skillProfeciencies, SkillProfeciency{"S", "Student"})
  skillProfeciencies = append(skillProfeciencies, SkillProfeciency{"J", "Junior"})
  skillProfeciencies = append(skillProfeciencies, SkillProfeciency{"E", "Experienced"})
  skillProfeciencies = append(skillProfeciencies, SkillProfeciency{"M", "Manager"})

}

func loadSkillProfeciency(value string) SkillProfeciency {
  loadSkillProfeciencies()

  for _, sp := range skillProfeciencies {
    if (value == sp.Value) {
      return sp.clone()
    }
  }
  return getEmptySkillProfeciency()
}

func parseSkillProfeciency(data interface{}) (SkillProfeciency, error) {

  print("start parsing a skill prof\n")

  value, ok := data.(string)
  if (ok) {
    return loadSkillProfeciency(value), nil
  }

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    print("looking for a 'map[string]interface{}' to parse a 'SkillProfeciency', but not found.\n")
    return getEmptySkillProfeciency(), errors.New("looking for a 'map[string]interface{}' to parse a 'SkillProfeciency', but not found.\n")
  }

  value = util.ParseString(dataMap, "value")
  print("value : " + value + "\n")
  return loadSkillProfeciency(value), nil

}

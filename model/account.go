package model

import (
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type Account struct {
  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City City `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`
  IsStudent bool `json:isstudent`

  Languages []AccountLanguage `json:languages`

  Educations []AccountEducation `json:educations`

  Skills []AccountSkill `json:skills`
}

func (acc Account) save(session *util.Session) error {
  if acc.Id > 0 {
    return acc.saveUpdate(session)
  }
  return acc.saveNew(session)
}

func (acc Account) saveUpdate(session *util.Session) error {
  return errors.New("account.saveUpdate is not implemented")
}

func (acc Account) saveNew(session *util.Session) error {
  query :=
    "INSERT INTO Account" +
    "(Name, Email, cityID, Phone, Password, isStudent) " +
    "VALUES($1, $2, $3, $4, $5, $6) " +
    "returning ID"

  id, err := util.Insert(query, acc.Name, acc.Email, acc.City.Id, acc.Phone, acc.Password, acc.IsStudent)

  if err != nil {
    return err
  }
  acc.Id = id
  session.PutAccountID(id)

  for _ , language := range acc.Languages {
    language.save(session)
  }

  for _ , education := range acc.Educations {
    education.save(session)
  }

  for _ , skill := range acc.Skills {
    skill.save(session)
  }

  return nil
}

func parseAccount(dataMap map[string]interface{}) (Account, error) {

  id := 0
  if dataMap["id"] != nil {
    id = int(dataMap["id"].(float64))
  }

  result := Account{Id: id}

  result.Name = util.ParseString(dataMap, "name")
  result.Email = util.ParseString(dataMap, "email")
  result.Phone = util.ParseString(dataMap, "phone")
  result.Password = util.ParseString(dataMap, "password")
  result.IsStudent = util.ParseBool(dataMap, "isstudent")

  cityMap := dataMap["city"].(map[string]interface{})
  city, _ := parseCity(cityMap)
  result.City = city

  langsArr := dataMap["languages"].([]interface{})
  result.Languages = parseAccountLanguages(langsArr)

  eduArr := dataMap["educations"].([]interface{})
  result.Educations = parseAccountEducations(eduArr)

  skillArr := dataMap["skills"].([]interface{})
  result.Skills = parseAccountSkills(skillArr)

  return result, nil
}

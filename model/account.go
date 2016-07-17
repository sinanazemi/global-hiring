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

  Languages []Language `json:languages`

  Educations []Education `json:educations`
}

func (acc Account) save() error {
  if acc.Id > 0 {
    return acc.saveUpdate()
  }
  return acc.saveNew()
}

func (acc Account) saveUpdate() error {
  return errors.New("account.saveUpdate is not implemented")
}

func (acc Account) saveNew() error {
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

  for _ , language := range acc.Languages {
    language.save(acc)
  }

  for _ , education := range acc.Educations {
    education.save(acc)
  }

  return nil
}

func parseAccount(dataMap map[string]interface{}) (interface{}, error) {

  print(dataMap)

  id := 0
  if dataMap["id"] != nil {
    id = int(dataMap["id"].(float64))
  }

  result := Account{Id: id}

  result.Name = dataMap["name"].(string)
  result.Email = dataMap["email"].(string)
  result.Phone = dataMap["phone"].(string)
  result.Password = dataMap["password"].(string)
  result.IsStudent = dataMap["isstudent"].(bool)

  cityMap := dataMap["city"].(map[string]interface{})
  city := parseCityReturn(cityMap)
  result.City = city

  langsArr := dataMap["languages"].([]interface{})
  result.Languages = parseLanguagesReturn(langsArr)

  eduArr := dataMap["educations"].([]interface{})
  result.Educations = parseEducationsReturn(eduArr)

  return result, nil
}

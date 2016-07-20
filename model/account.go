package model

import (
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Account struct {
  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City City `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`
  IsStudent bool `json:"isstudent"`

  Languages []AccountLanguage `json:"languages"`

  Educations []AccountEducation `json:"educations"`

  Skills []AccountSkill `json:"skills"`

  Certificates []AccountCertificate `json:"certificates"`
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

  if dataMap["certificates"] != nil {
    cerArr := dataMap["certificates"].([]interface{})
    result.Certificates = parseAccountCertificates(cerArr)
  }

  return result, nil
}

func loadAccount(session *util.Session) (Account, error) {
  query := "select ID, Name, Email, Phone, Password, IsStudent, cityID from Account Where ID = $1"
  accArr, _ := util.Select(readAccount, query, session.GetAccountID())
  account := accArr[0].(Account)

  account.City = loadCity(account.City.Id)

  account.Languages, _ = loadAccountLanguages(session)
  account.Educations, _ = loadAccountEducations(session)
  account.Skills, _ = loadAccountSkills(session)
  account.Certificates, _ = loadAccountCertificates(session)

  return account, nil
}

func readAccount(rows *sql.Rows) (interface{}, error) {
  var acc Account = Account{}
  err := rows.Scan(&acc.Id, &acc.Name, &acc.Email, &acc.Phone, &acc.Password, &acc.IsStudent, &acc.City.Id)

  return acc, err
}

func (acc Account) save(session *util.Session) error {
  if acc.Id <= 0 {
    return acc.saveNew(session)
  }
  return acc.saveUpdate(session)
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

  for _ , certificate := range acc.Certificates {
    certificate.save(session)
  }

  return nil
}

func (acc Account) saveUpdate(session *util.Session) error {
  return errors.New("account.saveUpdate is not implemented")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  account, err := loadAccount(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while loading account", http.StatusBadRequest}
  }
  account.Password = "" // :D
  return account, nil
}

func SaveAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  accountMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, _ := parseAccount(accountMap)

  err = account.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account", http.StatusBadRequest}
  }

  return account, nil
}

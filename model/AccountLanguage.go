package model

import (
  "errors"
  "strings"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

const LANGUAGE_PROFECIENCY_ELEMENTARY string = "E"
const LANGUAGE_PROFECIENCY_BASIC string = "B"
const LANGUAGE_PROFECIENCY_CONVERSATIONAL string = "C"
const LANGUAGE_PROFECIENCY_FLUENT string = "F"
const LANGUAGE_PROFECIENCY_NATIVE string = "N"

type AccountLanguage struct{
  Id     int    `json:"id"`
  Name string `json:"name"`
  Profeciency string `json:"profeciency"`
}

func parseAccountLanguage(dataMap map[string]interface{}) (AccountLanguage, error) {
  result := AccountLanguage{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Profeciency = util.ParseString(dataMap, "profeciency", " ")

  return result, nil
}

func parseAccountLanguages(langsArr []interface{}) []AccountLanguage {
  result := make([]AccountLanguage, 0)

  for _ , lang := range langsArr {
    lmap := lang.(map[string]interface{})
    lang, err := parseAccountLanguage(lmap)
    if err == nil {
      result = append(result, lang)
    }
  }
  return result
}

func (lang AccountLanguage) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountLanguage", "AccountID", lang.Id)
}

func (lang AccountLanguage) dataValidation(session *util.Session) error {
  errStr := ""

  if len(strings.TrimSpace(lang.Name)) <= 0 {
    errStr = errStr + "AccountLanguage.Name is Empty\n"
  }

  profeciencyCheck := false
  profeciencyCheck = profeciencyCheck || (lang.Profeciency == LANGUAGE_PROFECIENCY_ELEMENTARY)
  profeciencyCheck = profeciencyCheck || (lang.Profeciency == LANGUAGE_PROFECIENCY_BASIC)
  profeciencyCheck = profeciencyCheck || (lang.Profeciency == LANGUAGE_PROFECIENCY_CONVERSATIONAL)
  profeciencyCheck = profeciencyCheck || (lang.Profeciency == LANGUAGE_PROFECIENCY_FLUENT)
  profeciencyCheck = profeciencyCheck || (lang.Profeciency == LANGUAGE_PROFECIENCY_NATIVE)

  if (!profeciencyCheck) {
    errStr = errStr + "AccountLanguage.Profeciency is not valid\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (lang AccountLanguage) insertValidation(session *util.Session) error {
  return lang.dataValidation(session)
}

func (lang AccountLanguage) updateValidation(session *util.Session) error {
  err := lang.accountValidation(session)
  if err != nil{
    return err
  }
  return lang.dataValidation(session)
}

func (lang AccountLanguage) deleteValidation(session *util.Session) error {
  return lang.accountValidation(session)
}

func LoadAccountLanguages(session *util.Session) ([]AccountLanguage, error) {
  query :=
    "SELECT ID, Name, Profeciency " +
    "FROM AccountLanguage " +
    " WHERE AccountID = $1";

    var result = make([]AccountLanguage, 0)

    languages, err := util.Select(readAccountLanguage, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyLang := range languages {
      language, _ := dummyLang.(AccountLanguage)
      result = append(result, language)
    }

    return result, nil
}

func readAccountLanguage(rows *sql.Rows) (interface{}, error) {

    var lang AccountLanguage = AccountLanguage{}
    err := rows.Scan(&lang.Id, &lang.Name, &lang.Profeciency)

    return lang, err
}

func (lang AccountLanguage) save(session *util.Session) error {
  if lang.Id <= 0 {
    return lang.saveNew(session)
  }
  return lang.saveUpdate(session)
}

func (lang AccountLanguage) saveNew(session *util.Session) error {

  err := lang.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountLanguage" +
    "(Name, Profeciency, AccountID) " +
    "VALUES($1, $2, $3) " +
    "returning ID"

  id, err := util.Insert(query, lang.Name, lang.Profeciency, session.GetAccountID())

  if err != nil {
    return err
  }

  lang.Id = id
  return nil
}

func (lang AccountLanguage) saveUpdate(session *util.Session) error {
  err := lang.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountLanguage " +
    "SET " +
    "Name = $1, " +
    "Profeciency = $2 " +
    "WHERE ID = $3 "

  err = util.Update(query, lang.Name, lang.Profeciency, lang.Id)

  return err

}

func (lang AccountLanguage) delete(session *util.Session) error {
  err := lang.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountLanguage " +
    "WHERE ID = $1 "

  err = util.Update(query, lang.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveLanguage(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  langMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Language", http.StatusBadRequest}
  }

  lang, _ := parseAccountLanguage(langMap)

  err = lang.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Language", http.StatusBadRequest}
  }

  return lang, nil
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  langMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Language", http.StatusBadRequest}
  }

  lang, _ := parseAccountLanguage(langMap)

  err = lang.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Language", http.StatusBadRequest}
  }

  return lang, nil
}

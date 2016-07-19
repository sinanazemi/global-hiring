package model

import (
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

func (lang AccountLanguage) save(session *util.Session) error {
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

package model

import (
  "github.com/sinanazemi/global-hiring/util"
)

const LANGUAGE_PROFECIENCY_ELEMENTARY string = "E"
const LANGUAGE_PROFECIENCY_BASIC string = "B"
const LANGUAGE_PROFECIENCY_CONVERSATIONAL string = "C"
const LANGUAGE_PROFECIENCY_FLUENT string = "F"
const LANGUAGE_PROFECIENCY_NATIVE string = "N"

type Language struct{
  Id     int    `json:"id"`
  Name string `json:"name"`
  Profeciency string `json:"profeciency"`
}

func parseLanguage(dataMap map[string]interface{}) (Language, error) {
  result := Language{}

  result.Name = dataMap["name"].(string)
  result.Profeciency = dataMap["profeciency"].(string)

  return result, nil
}

func parseLanguages(langsArr []interface{}) []Language {
  result := make([]Language, 0)

  for _ , lang := range langsArr {
    lmap := lang.(map[string]interface{})
    lang, err := parseLanguage(lmap)
    if err == nil {
      result = append(result, lang)
    }
  }
  return result
}

func (lang Language) save(account Account) error {
  query :=
    "INSERT INTO AccountLanguage" +
    "(Name, Profeciency, AccountID) " +
    "VALUES($1, $2, $3) " +
    "returning ID"

  id, err := util.Insert(query, lang.Name, lang.Profeciency, account.Id)

  if err != nil {
    return err
  }

  lang.Id = id
  return nil
}

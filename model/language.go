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

func parseLanguage(dataMap map[string]interface{}) (interface{}, error) {
  result := Language{}

  result.Name = dataMap["name"].(string)
  result.Profeciency = dataMap["profeciency"].(string)

  return result, nil
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

package model

import (
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

func GetCities(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getCities(), nil
}

func GetMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getMainServices(), nil
}

func GetSkills(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  id := util.GetID(r)
  return getSkills(id), nil
}

func GetDegrees(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getDegrees(), nil
}

func SaveAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {
  accountMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, _ := parseAccount(accountMap)

  err = account.save()
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account", http.StatusBadRequest}
  }

  return account, nil
}

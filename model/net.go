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

func SaveSkill(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  skillMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Skill", http.StatusBadRequest}
  }

  acs, _ := parseAccountSkill(skillMap)

  err = acs.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving skill", http.StatusBadRequest}
  }

  return acs, nil
}

func DeleteSkill(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  skillMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Skill", http.StatusBadRequest}
  }

  acs, _ := parseAccountSkill(skillMap)

  err = acs.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting skill", http.StatusBadRequest}
  }

  return acs, nil
}

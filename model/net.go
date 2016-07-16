package model

import (
  "errors"
  "strconv"
  "encoding/json"
	"io/ioutil"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type idStruct struct {
  Id int `json:"id"`
}

func parseIdRequest(r *http.Request) int {

	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
    print("Could not read request\n")
		return -1
	}

	var idJson idStruct
	e = json.Unmarshal(data, &idJson)
	if e != nil {
    print("Could not parse JSON\n")
		return -1
	}
	return idJson.Id
}

type JsonParser func(map[string]interface{}) (interface{}, error)

func parseJsonRequest(r *http.Request, parse JsonParser) (interface{}, error) {

	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
    print("Could not read request")
		return nil, errors.New("Could not read request")
	}

	// turn the request body (JSON) into a book object
	var dataMap map[string]interface{}
	e = json.Unmarshal(data, &dataMap)
	if e != nil {
    print("Could not parse JSON")
		return nil, errors.New("Could not parse JSON")
	}

  return parse(dataMap)
}

func GetCities(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getCities(), nil
}

func GetMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getMainServices(), nil
}

func GetSkills(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  param := r.URL.Query().Get("id")
  print(param + "\n")

  id, e := strconv.Atoi(param)
  if e != nil {
    return nil, &util.HandlerError{e, "Id should be an integer", http.StatusBadRequest}
  }

  return getSkills(id), nil
}

func SaveAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {
  accountJ, err := parseJsonRequest(r, parseAccount)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account := accountJ.(Account)

  err = account.save()
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account", http.StatusBadRequest}
  }

  return account, nil
}

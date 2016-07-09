package model

import (
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

func GetMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getMainServices(), nil
}

func GetSkills(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  //param := mux.Vars(r)["id"]
  param := r.URL.Query().Get("id")
  print(param + "\n")

  id, e := strconv.Atoi(param)
  if e != nil {
    return nil, &util.HandlerError{e, "Id should be an integer", http.StatusBadRequest}
  }

//  id := parseIdRequest(r)


  return getSkills(id), nil
}

/*func parseMainServiceRequest(r *http.Request) (MainService, *util.HandlerError) {

	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
    print("Could not read request")
		return MainService{}, &util.HandlerError{e, "Could not read request", http.StatusBadRequest}
	}

	// turn the request body (JSON) into a book object
	var service MainService
	e = json.Unmarshal(data, &service)
	if e != nil {
    print("Could not parse JSON")
		return MainService{}, &util.HandlerError{e, "Could not parse JSON", http.StatusBadRequest}
	}

  print(service.Id)
  return service, nil
}*/

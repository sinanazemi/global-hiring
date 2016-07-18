package util

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "errors"
)

func ParseJsonRequest(r *http.Request) (map[string]interface{}, error) {

	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, errors.New("Could not read request")
	}

	// turn the request body (JSON) into a book object
	var result map[string]interface{}
	e = json.Unmarshal(data, &result)
	if e != nil {
		return nil, errors.New("Could not parse JSON")
	}

  return result, nil
}

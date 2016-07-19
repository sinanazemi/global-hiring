package util

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "errors"
  "log"
  "strings"
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

func ParseString(dataMap map[string]interface{}, paramName string, defaultValue ...string) string {

  defaultVal := ""
  if (defaultValue != nil && len(defaultValue) > 0) {
    defaultVal = defaultValue[0]
  }

  if dataMap[paramName] == nil {
    log.Printf("No Prameter is set as '%s', default value '%s' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  str, ok := dataMap[paramName].(string)
  if (! ok) {
    log.Printf("Prameter '%s' is not a string! default value '%s' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  return str
}

func ParseInteger(dataMap map[string]interface{}, paramName string, defaultValue ...int) int {

  defaultVal := -1
  if (defaultValue != nil && len(defaultValue) > 0) {
    defaultVal = defaultValue[0]
  }

  if dataMap[paramName] == nil {
    log.Printf("No Prameter is set as '%s', default value '%d' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  i, ok := dataMap[paramName].(float64)
  if (! ok) {
    log.Printf("Prameter '%s' is not a numeric! default value '%d' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  return int(i)
}

func ParseFloat(dataMap map[string]interface{}, paramName string, defaultValue ...float64) float64 {

  defaultVal := -1.0
  if (defaultValue != nil && len(defaultValue) > 0) {
    defaultVal = defaultValue[0]
  }

  if dataMap[paramName] == nil {
    log.Printf("No Prameter is set as '%s', default value '%-6.2f' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  f, ok := dataMap[paramName].(float64)
  if (! ok) {
    log.Printf("Prameter '%s' is not a numeric! default value '%-6.2f' will be returned.", paramName, defaultVal)
    return defaultVal
  }
  return f
}

func ParseBool(dataMap map[string]interface{}, paramName string, defaultValue ...bool) bool {

  defaultVal := false
  if (defaultValue != nil && len(defaultValue) > 0) {
    defaultVal = defaultValue[0]
  }

  if dataMap[paramName] == nil {
    log.Printf("No Prameter is set as '%s', default value '%t' will be returned.", paramName, defaultVal)
    return defaultVal
  }

  b, ok := dataMap[paramName].(bool)
  if (ok) {
    return b
  }

  bStr, ok := dataMap[paramName].(string)
  if (ok) {
    if (strings.ToLower(bStr) == "false"){
      return false
    }
    if (strings.ToLower(bStr) == "true"){
      return true
    }
  }

  log.Printf("Prameter '%s' is not a boolean! default value '%t' will be returned.", paramName, defaultVal)
  return defaultVal
}

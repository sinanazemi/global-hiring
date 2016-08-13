package util

import (
  "encoding/json"
	"fmt"
	"log"
  "strconv"

	"net/http"
)

// error response contains everything we need to use http.Error
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

// a custom type that we can use for handling errors and formatting responses
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to
  log.Printf("ServeHTTP %s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)

  session, er := GetSession(w, r)
  if er == nil {
      session.Refresh()
  }

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("finished %s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

type idStruct struct {
  Id int `json:"id"`
}

func GetID(r *http.Request, paramNames ...string) int {

  query := r.URL.Query();

  paramNames = append(paramNames, "id")

  for _ , paramName := range paramNames {
    param := query.Get(paramName)
    id, e := strconv.Atoi(param)
    if e == nil {
      return id
    }
  }
  return -1
}

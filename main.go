package main

import (
  "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
//  "global-hiring/mapper"

//	"github.com/gorilla/mux"

)

// error response contains everything we need to use http.Error
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// book model
type user struct {
  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City string `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`
}

// book model
type firstName struct {
	Name  string `json:"name"`
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to
	log.Print("ServeHTTP")

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
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

func parseUserRequest(r *http.Request) (user, *handlerError) {
	// the book payload is in the request body
	log.Print("parseUserRequest")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return user{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}

	// turn the request body (JSON) into a book object
	var payload user
	e = json.Unmarshal(data, &payload)
	if e != nil {
		return user{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	}

	return payload, nil
}

func getFirstName(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	log.Print("getFirstName")
  var f = firstName{"John Doe"}
  log.Print(f.Name)
  return f, nil
}

func getMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

  return nil,nil;
}

func getSkills(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

  return nil,nil;
}

func saveAccount()  (w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

  return nil,nil;
}



func main() {

  // command line flags
	port := flag.Int("port", 80, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)
  //http.Handle("/firstName", handler(getFirstName))//.Methods("POST")
  http.Handle("/firstName/{id}", handler(getFirstName))//.Methods("POST")
  http.Handle("/mainService", handler(getMainServices))
  http.Handle("/skills/{id}", handler(getSkills))//.Methods("POST")


	log.Printf("Running on port %d\n", *port)


	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())


}

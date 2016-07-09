package main

import (
  "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

  "github.com/sinanazemi/global-hiring/model"
  "github.com/sinanazemi/global-hiring/util"

)

// book model
type user struct {
  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City string `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`
}

func parseUserRequest(r *http.Request) (user, *util.HandlerError) {
	// the book payload is in the request body
	log.Print("parseUserRequest")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return user{}, &util.HandlerError{e, "Could not read request", http.StatusBadRequest}
	}

	// turn the request body (JSON) into a book object
	var payload user
	e = json.Unmarshal(data, &payload)
	if e != nil {
		return user{}, &util.HandlerError{e, "Could not parse JSON", http.StatusBadRequest}
	}

	return payload, nil
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

  http.Handle("/mainService", util.Handler(model.GetMainServices))
  //http.HandleFunc("/mainService", model.GetMainServices)
  //http.Handle("/skills/{id}", util.Handler(model.GetSkills))
  http.Handle("/skills", util.Handler(model.GetSkills))


	log.Printf("Running on port %d\n", *port)


	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())


}

package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"

  "github.com/sinanazemi/global-hiring/model"
  "github.com/sinanazemi/global-hiring/util"

)

func main() {

  // command line flags
	port := flag.Int("port", 80, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)

	http.Handle("/cities", util.Handler(model.GetCities))
	http.Handle("/mainServices", util.Handler(model.GetMainServices))
	http.Handle("/skills", util.Handler(model.GetSkills))
	http.Handle("/saveAccount", util.Handler(model.SaveAccount))

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())


}

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
	http.Handle("/degrees", util.Handler(model.GetDegrees))
	http.Handle("/mainServices", util.Handler(model.GetMainServices))
	http.Handle("/skills", util.Handler(model.GetSkills))
	http.Handle("/volunteeringCauses", util.Handler(model.GetVolunteeringCauses))


	// Editting Skills
	http.Handle("/account", util.Handler(model.GetAccount))
	http.Handle("/saveAccount", util.Handler(model.SaveAccount))

	// Editting Skills
	http.Handle("/saveSkill", util.Handler(model.SaveSkill))
	http.Handle("/deleteSkill", util.Handler(model.DeleteSkill))

	// Editting Certificates
	http.Handle("/saveCertificate", util.Handler(model.SaveCertificate))
	http.Handle("/deleteCertificate", util.Handler(model.DeleteCertificate))

	// Editting Educations
	http.Handle("/saveEducation", util.Handler(model.SaveEducation))
	http.Handle("/deleteEducation", util.Handler(model.DeleteEducation))

	// Editting Languages
	http.Handle("/saveLanguage", util.Handler(model.SaveLanguage))
	http.Handle("/deleteLanguage", util.Handler(model.DeleteLanguage))

	// Editting Work Histories
	http.Handle("/saveWork", util.Handler(model.SaveWork))
	http.Handle("/deleteWork", util.Handler(model.DeleteWork))

	// Editting Volunteering
	http.Handle("/saveVolunteering", util.Handler(model.SaveVolunteering))
	http.Handle("/deleteVolunteering", util.Handler(model.DeleteVolunteering))

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}

package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"

	"github.com/sinanazemi/global-hiring/util"
  "github.com/sinanazemi/global-hiring/model"

)


func Authenticate(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {
	_ , err := util.Authenticate(w, r)
	if (err != nil) {
		return nil, err
	}

  return model.GetAccount(w, r)
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

	http.Handle("/redirect", util.Handler(util.RedirectCheck))
	http.Handle("/authenticate", util.Handler(Authenticate))
	http.Handle("/logout", util.Handler(util.Logout))

	http.Handle("/cities", util.Handler(model.GetCities))
	http.Handle("/degrees", util.Handler(model.GetDegrees))
	http.Handle("/mainServices", util.Handler(model.GetMainServices))
	http.Handle("/skills", util.Handler(model.GetSkills))
	http.Handle("/volunteeringCauses", util.Handler(model.GetVolunteeringCauses))
	http.Handle("/occupations", util.Handler(model.GetOccupations))


	// Editting Accounts
	http.Handle("/account", util.Handler(model.GetAccount))
	http.Handle("/accountStrength", util.Handler(model.GetAccountStrength))
	http.Handle("/saveAccount", util.Handler(model.SaveAccount))
	http.Handle("/completeAccount", util.Handler(model.CompleteAccount))

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

	// Editting Honors
	http.Handle("/saveHonor", util.Handler(model.SaveHonor))
	http.Handle("/deleteHonor", util.Handler(model.DeleteHonor))

	// Editting Courses
	http.Handle("/saveCourse", util.Handler(model.SaveCourse))
	http.Handle("/deleteCourse", util.Handler(model.DeleteCourse))

	// Editting Tests
	http.Handle("/saveTest", util.Handler(model.SaveTest))
	http.Handle("/deleteTest", util.Handler(model.DeleteTest))

	// Editting Projects
	http.Handle("/saveProject", util.Handler(model.SaveProject))
	http.Handle("/deleteProject", util.Handler(model.DeleteProject))

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}

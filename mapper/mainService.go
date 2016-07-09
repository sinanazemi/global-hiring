package model

import (
  net/http
)

type MainService struct{
  Name string `json="name"`
  Skills []Skill `json="skills"`
}

func getMainServices() []MainService {

  // list of all of the books
	var services = make([]MainService, 0)

  // bootstrap some data
  services = append(services, MainService{"Service 1", nil})
  services = append(services, MainService{"Service 2", nil})
  services = append(services, MainService{"Service 3", nil})
  services = append(services, MainService{"Service 4", nil})

  return services

}

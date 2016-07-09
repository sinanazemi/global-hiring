package model

import(
  //"encoding/json"
)

type MainService struct{
  Id     int    `json:"id"`
  Name string `json:"name"`
  Skills []Skill `json:"skills"`
}

func getMainServices() []MainService {

  // list of all of services
	var services = make([]MainService, 0)

  // bootstrap some data
  services = append(services, MainService{1, "Service 1", nil})
  services = append(services, MainService{2, "Service 2", nil})
  services = append(services, MainService{3, "Service 3", nil})
  services = append(services, MainService{4, "Service 4", nil})

  return services

}

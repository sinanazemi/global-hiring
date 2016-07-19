package model

import(
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type MainService struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
  Question string `json:"question"`
  IsSelected bool `json:"isselected"`
  Skills []Skill `json:"skills"`
}

func getMainServices() []MainService {

  // list of all of services
	var result = make([]MainService, 0)

  services, err := util.Select(readMainService, "select * from MainService")

  if err != nil {
    return result
  }

  for _, dummyService := range services {

    service, ok := dummyService.(MainService)
    if !ok {
        // service was not of type MainService. The assertion failed
        return make([]MainService, 0)
    }
    // service is of type MainService
    service.Skills = getSkills(service.Id)
    result = append(result, service)
  }
  return result

}

func readMainService(rows *sql.Rows) (interface{}, error) {

  var service MainService = MainService{}
  err := rows.Scan(&service.Id, &service.Name, &service.Question)
  service.IsSelected = false

  return service, err
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getMainServices(), nil
}

package model

import(
  "database/sql"
  "github.com/sinanazemi/global-hiring/util"
)

type MainService struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
  Skills []Skill `json:"skills"`
}

func getMainServices() []MainService {

  // list of all of services
	var result = make([]MainService, 0)

  services, err := util.Select(readMainService, "select * from public.\"MainService\"")

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
    result = append(result, service)
  }
  return result

}

func readMainService(rows *sql.Rows) (interface{}, error) {

  var service MainService = MainService{-1, "", nil}
  err := rows.Scan(&service.Id, &service.Name)

  return service, err
}

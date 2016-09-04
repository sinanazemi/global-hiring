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

  UnselectImageURL string `json:"unselectimageurl"`
  SelectImageURL string `json:"selectimageurl"`
}

func getMainServices() []MainService {

  // list of all of services
	var result = make([]MainService, 0)

  query :=
    "select m.*, COALESCE(i1.ID, 0), COALESCE(i2.ID, 0) from MainService m " +
    "left outer join Image i1 on i1.category = 'MS1u' and i1.ParentID = m.ID " +
    "left outer join Image i2 on i2.category = 'MS1s' and i2.ParentID = m.ID "


  services, err := util.Select(readMainService, query)

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

  var idu, ids int

  var service MainService = MainService{}
  err := rows.Scan(&service.Id, &service.Name, &service.Question, &idu, &ids)
  service.IsSelected = false

  service.UnselectImageURL = util.GetImageURL(idu)
  service.SelectImageURL = util.GetImageURL(ids)

  return service, err
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetMainServices(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getMainServices(), nil
}

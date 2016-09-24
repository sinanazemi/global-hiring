package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type MainService struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
  Question string `json:"question"`
  IsSelected bool `json:"isselected"` // dummy field, just for leva
  Skills []Skill `json:"skills"`

  UnselectImageURL string `json:"unselectimageurl"`
  SelectImageURL string `json:"selectimageurl"`
}

func parseMainService(data interface{}) (MainService, error) {
  result := MainService{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'MainService', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Question = util.ParseString(dataMap, "question")
  result.IsSelected = util.ParseBool(dataMap, "isselected")

  result.Skills = parseSkills(dataMap["skills"])

  result.UnselectImageURL = util.ParseString(dataMap, "unselectimageurl")
  result.SelectImageURL = util.ParseString(dataMap, "selectimageurl")

  return result, nil
}

func parseMainServices(data interface{}) []MainService {
  result := make([]MainService, 0)

  mArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a '[]interface{}' to parse an array of 'MainService's, but not found.\n")
    return result
  }

  for _ , m := range mArr {
    mainService, err := parseMainService(m)
    if (err == nil) {
      result = append(result, mainService)
    }
  }
  return result
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

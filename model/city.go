package model

import(
  "database/sql"
  "github.com/sinanazemi/global-hiring/util"
)

type City struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func getCities() []City {

  // list of all of services
	var result = make([]City, 0)

  cities, err := util.Select(readCity, "select * from City")

  if err != nil {
    return result
  }

  for _, dummyCity := range cities {

    city, ok := dummyCity.(City)
    if !ok {
        // service was not of type MainService. The assertion failed
        return make([]City, 0)
    }
    // service is of type MainService
    result = append(result, city)
  }
  return result

}

func readCity(rows *sql.Rows) (interface{}, error) {

  var city City = City{}
  err := rows.Scan(&city.Id, &city.Name)

  return city, err
}

func parseCity(dataMap map[string]interface{}) (interface{}, error) {

  result := City{}

  result.Id = int(dataMap["id"].(float64))
  result.Name = dataMap["name"].(string)

  return result, nil
}

func parseCityReturn(cityMap map[string]interface{}) City {

  cityParsed, _ := parseCity(cityMap)
  city := cityParsed.(City)
  return city
}

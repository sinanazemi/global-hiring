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
    city, _ := dummyCity.(City)
    result = append(result, city)
  }

  return result
}

func readCity(rows *sql.Rows) (interface{}, error) {

  var city City = City{}
  err := rows.Scan(&city.Id, &city.Name)

  return city, err
}

func parseCity(dataMap map[string]interface{}) (City, error) {

  result := City{}

  result.Id = int(dataMap["id"].(float64))
  result.Name = dataMap["name"].(string)

  return result, nil
}

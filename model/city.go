package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type City struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func parseCity(data interface{}) (City, error) {

  id, ok := data.(float64)
  if (ok) {
    return loadCity(int(id)), nil
  }

  result := City{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'City', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}

func loadCities(where string, args ...interface{}) []City {

	var result = make([]City, 0)

  cities, err := util.Select(readCity, "select * from City where " + where, args...)

  if err != nil {
    return result
  }

  for _, dummyCity := range cities {
    city, _ := dummyCity.(City)
    result = append(result, city)
  }

  return result
}

func loadCity(cityID int) City {

	return loadCities("ID = $1", cityID)[0]
}

func getCities() []City {

	return loadCities("1=1")
}

func readCity(rows *sql.Rows) (interface{}, error) {

  var city City = City{}
  err := rows.Scan(&city.Id, &city.Name)

  return city, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetCities(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getCities(), nil
}

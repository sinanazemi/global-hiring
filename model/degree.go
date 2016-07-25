package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Degree struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func parseDegree(data interface{}) (Degree, error) {

  id, ok := data.(float64)
  if (ok) {
    return loadDegree(int(id)), nil
  }

  result := Degree{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'Degree', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}

func loadDegrees(where string, args ...interface{}) []Degree {

  // list of all of services
	var result = make([]Degree, 0)

  degrees, err := util.Select(readDegree, "select * from Degree where " + where, args...)

  if err != nil {
    return result
  }

  for _, dummyDegree := range degrees {
    degree, _ := dummyDegree.(Degree)
    result = append(result, degree)
  }

  return result
}

func loadDegree(degreeID int) Degree {

  return loadDegrees("ID = $1", degreeID)[0]
}

func getDegrees() []Degree {

  return loadDegrees("1=1")
}

func readDegree(rows *sql.Rows) (interface{}, error) {

  var degree Degree = Degree{}
  err := rows.Scan(&degree.Id, &degree.Name)

  return degree, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetDegrees(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getDegrees(), nil
}

package model

import(
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Degree struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func parseDegree(dataMap map[string]interface{}) (Degree, error) {

  result := Degree{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}

func LoadDegrees(where string, args ...interface{}) []Degree {

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

func LoadDegree(degreeID int) Degree {

  return LoadDegrees("ID = $1", degreeID)[0]
}

func getDegrees() []Degree {

  return LoadDegrees("1=1")
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

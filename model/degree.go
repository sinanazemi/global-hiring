package model

import(
  "database/sql"
  "github.com/sinanazemi/global-hiring/util"
)

type Degree struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func getDegrees() []Degree {

  // list of all of services
	var result = make([]Degree, 0)

  degrees, err := util.Select(readDegree, "select * from Degree")

  if err != nil {
    return result
  }

  for _, dummyDegree := range degrees {
    degree, _ := dummyDegree.(Degree)
    result = append(result, degree)
  }

  return result
}

func readDegree(rows *sql.Rows) (interface{}, error) {

  var degree Degree = Degree{}
  err := rows.Scan(&degree.Id, &degree.Name)

  return degree, err
}

func parseDegree(dataMap map[string]interface{}) (Degree, error) {

  result := Degree{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}

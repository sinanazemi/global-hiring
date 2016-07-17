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

    degree, ok := dummyDegree.(Degree)
    if !ok {
        // service was not of type MainService. The assertion failed
        return make([]Degree, 0)
    }
    // service is of type MainService
    result = append(result, degree)
  }
  return result

}

func readDegree(rows *sql.Rows) (interface{}, error) {

  var degree Degree = Degree{}
  err := rows.Scan(&degree.Id, &degree.Name)

  return degree, err
}

func parseDegree(dataMap map[string]interface{}) (interface{}, error) {

  result := Degree{}

  result.Id = int(dataMap["id"].(float64))
  result.Name = dataMap["name"].(string)

  return result, nil
}

func parseDegreeReturn(degreeMap map[string]interface{}) Degree {

  degreeParsed, _ := parseDegree(degreeMap)
  degree := degreeParsed.(Degree)
  return degree
}

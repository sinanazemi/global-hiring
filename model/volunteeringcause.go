package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type VolunteeringCause struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func parseVolunteeringCause(data interface{}) (VolunteeringCause, error) {

  id, ok := data.(float64)
  if (ok) {
    return loadVolunteeringCause(int(id)), nil
  }

  result := VolunteeringCause{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'VolunteeringCause', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}


func loadVolunteeringCauses(where string, args ...interface{}) []VolunteeringCause {

  // list of all of services
	var result = make([]VolunteeringCause, 0)

  causes, err := util.Select(readDegree, "select * from VolunteeringCause where " + where, args...)

  if err != nil {
    return result
  }

  for _, dummyCause := range causes {
    cause, _ := dummyCause.(VolunteeringCause)
    result = append(result, cause)
  }

  return result
}

func loadVolunteeringCause(causeID int) VolunteeringCause {

  return loadVolunteeringCauses("ID = $1", causeID)[0]
}

func getVolunteeringCauses() []VolunteeringCause {

  return loadVolunteeringCauses("1=1")
}

func readVolunteeringCause(rows *sql.Rows) (interface{}, error) {

  var cause VolunteeringCause = VolunteeringCause{}
  err := rows.Scan(&cause.Id, &cause.Name)

  return cause, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetVolunteeringCauses(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  return getVolunteeringCauses(), nil
}

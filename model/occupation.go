package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Occupation struct{
  Id     int  `json:"id"`
  Name string `json:"name"`
}

func parseOccupation(data interface{}) (Occupation, error) {

  id, ok := data.(float64)
  if (ok) {
    return loadOccupation(int(id)), nil
  }

  result := Occupation{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse a 'Occupation', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")

  return result, nil
}


func loadOccupations(where string, args ...interface{}) []Occupation {

  // list of all of services
	var result = make([]Occupation, 0)

  occupations, err := util.Select(readOccupation, "select id, name from Occupation where " + where, args...)

  if err != nil {
    return result
  }

  for _, dummyOcc := range occupations {
    occ, _ := dummyOcc.(Occupation)
    result = append(result, occ)
  }

  return result
}

func loadOccupation(occID int) Occupation {

  return loadOccupations("ID = $1", occID)[0]
}

func getOccupations(session *util.Session) []Occupation {

  return loadOccupations("accountID = $1", session.GetAccountID())
}

func readOccupation(rows *sql.Rows) (interface{}, error) {

  var occ Occupation = Occupation{}
  err := rows.Scan(&occ.Id, &occ.Name)

  return occ, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetOccupations(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  return getOccupations(session), nil
}

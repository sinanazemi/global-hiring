package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountVolunteering struct {
  Id     int  `json:"id"`
  Organization string `json:"organization"`
  Role string `json:"role"`
  Cause VolunteeringCause `json:"cause"`
  FromMonth Month `json:"frommonth"`
  FromYear int `json:"fromyear"`
  ToMonth Month `json:"tomonth"`
  ToYear int `json:"toyear"`
  Description string `json:"description"`
}

func parseAccountVolunteering(data interface{}) (AccountVolunteering, error) {
  result := AccountVolunteering{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountVolunteering', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.Organization = util.ParseString(dataMap, "organization")
  result.Role = util.ParseString(dataMap, "role")

  result.Cause, _ = parseVolunteeringCause(dataMap["cause"])

  result.FromMonth, _ = parseMonth(dataMap["frommonth"])
  result.FromYear = util.ParseInteger(dataMap, "fromyear")

  result.ToMonth, _ = parseMonth(dataMap["tomonth"])
  result.ToYear = util.ParseInteger(dataMap, "toyear")

  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountVolunteerings(data interface{}) []AccountVolunteering {
  result := make([]AccountVolunteering, 0)

  vArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a 'm[]interface{}' to parse an array of 'AccountVolunteering's, but not found.\n")
    return result
  }

  for _ , vol := range vArr {
    volunteering, err := parseAccountVolunteering(vol)
    if (err == nil) {
      result = append(result, volunteering)
    }
  }
  return result
}

func (avl *AccountVolunteering) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountVolunteering", "AccountID", avl.Id)
}

func (avl *AccountVolunteering) dataValidation(session *util.Session) error {

  errStr := ""

  if len(strings.TrimSpace(avl.Organization)) <= 0 {
    errStr = errStr + "AccountVolunteering.Organization is Empty\n"
  }

  if len(strings.TrimSpace(avl.Role)) <= 0 {
    errStr = errStr + "AccountVolunteering.Role is Empty\n"
  }

  if (avl.Cause.Id <= 0) {
    errStr = errStr + "Invalid Cause in AccountVolunteering\n"
  }

  if (avl.FromYear <= 0 || avl.FromMonth.isEmpty()) {
    errStr = errStr + "Invalid start date in AccountVolunteering\n"
  }

  if (avl.ToYear <= 0 || avl.ToMonth.isEmpty()) {
    errStr = errStr + "Invalid end date in AccountVolunteering\n"
  }

  var from = (avl.FromYear * 12) + avl.FromMonth.Value
  var to = (avl.ToYear * 12) + avl.ToMonth.Value

  if (to < from) {
    errStr = errStr + "Invalid relation between dates in AccountVolunteering\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (avl *AccountVolunteering) insertValidation(session *util.Session) error {
  return avl.dataValidation(session)
}

func (avl *AccountVolunteering) updateValidation(session *util.Session) error {
  err := avl.accountValidation(session)
  if err != nil{
    return err
  }
  return avl.dataValidation(session)
}

func (avl *AccountVolunteering) deleteValidation(session *util.Session) error {
  return avl.accountValidation(session)
}

func loadAccountVolunteerings(session *util.Session) ([]AccountVolunteering, error) {

  query :=
    "SELECT ID, Organization, Role, Cause, FromMonth, FromYear, ToMonth, ToYear, Description " +
    "FROM AccountVolunteering " +
    " WHERE AccountID = $1";

    var result = make([]AccountVolunteering, 0)

    vols, err := util.Select(readAccountVolunteering, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyVol := range vols {

      vol, _ := dummyVol.(AccountVolunteering)

      vol.Cause = loadVolunteeringCause(vol.Cause.Id)
      vol.FromMonth = loadMonth(vol.FromMonth.Value)
      vol.ToMonth = loadMonth(vol.ToMonth.Value)

      result = append(result, vol)
    }

    return result, nil
}

func readAccountVolunteering(rows *sql.Rows) (interface{}, error) {

    var vol AccountVolunteering = AccountVolunteering{}
    err := rows.Scan(&vol.Id, &vol.Organization, &vol.Role, &vol.Cause.Id, &vol.FromMonth.Value, &vol.FromYear, &vol.ToMonth.Value, &vol.ToYear, &vol.Description)

    return vol, err
}

func (avl *AccountVolunteering) save(session *util.Session) error {
  if avl.Id <= 0 {
    return avl.saveNew(session)
  }
  return avl.saveUpdate(session)
}

func (vol *AccountVolunteering) saveNew(session *util.Session) error {

  err := vol.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountVolunteering " +
    "(Organization, Role, Cause, FromMonth, FromYear, ToMonth, ToYear, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) " +
    "returning ID"

  id, err := util.Insert(query, vol.Organization, vol.Role, vol.Cause.Id, vol.FromMonth.Value, vol.FromYear, vol.ToMonth.Value, vol.ToYear, vol.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  vol.Id = id
  return nil
}

func (vol *AccountVolunteering) saveUpdate(session *util.Session) error {
  err := vol.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountVolunteering " +
    "SET " +
    "Organization = $1, " +
    "Role = $2, " +
    "Cause = $3, " +
    "FromMonth = $4, " +
    "FromYear = $5, " +
    "ToMonth = $6, " +
    "ToYear = $7, " +
    "Description = $8 " +
    "WHERE ID = $9 "

  err = util.Update(query, vol.Organization, vol.Role, vol.Cause.Id, vol.FromMonth.Value, vol.FromYear, vol.ToMonth.Value, vol.ToYear, vol.Description, vol.Id)

  return err

}

func (avl *AccountVolunteering) delete(session *util.Session) error {
  err := avl.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountVolunteering " +
    "WHERE ID = $1 "

  err = util.Update(query, avl.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveVolunteering(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  volMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Volunteering", http.StatusBadRequest}
  }

  vol, _ := parseAccountVolunteering(volMap)

  err = vol.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Volunteering", http.StatusBadRequest}
  }

  return vol, nil
}

func DeleteVolunteering(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  volMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Volunteering", http.StatusBadRequest}
  }

  vol, _ := parseAccountVolunteering(volMap)

  err = vol.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Volunteering", http.StatusBadRequest}
  }

  return vol, nil
}

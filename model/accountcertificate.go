package model

import (
  "strings"
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountCertificate struct {
  Id     int  `json:"id"`
  Name string `json:"name"`
  Authority string `json:"authority"`
  License string `json:"license"`
  Url string `json:"url"`
  Description string `json:"description"`
}

func parseAccountCertificate(dataMap map[string]interface{}) (AccountCertificate, error) {
  result := AccountCertificate{}

  result.Id = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.Authority = util.ParseString(dataMap, "authority")
  result.License = util.ParseString(dataMap, "license")
  result.Url = util.ParseString(dataMap, "url")
  result.Description = util.ParseString(dataMap, "description")

  return result, nil
}

func parseAccountCertificates(cerArr []interface{}) []AccountCertificate {
  result := make([]AccountCertificate, 0)

  for _ , cer := range cerArr {
    cmap := cer.(map[string]interface{})
    certificate, err := parseAccountCertificate(cmap)
    if (err == nil) {
      result = append(result, certificate)
    }
  }
  return result
}

func (cer AccountCertificate) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountCertificate", "AccountID", cer.Id)
}

func (cer AccountCertificate) dataValidation(session *util.Session) error {
  errStr := ""

  if len(strings.TrimSpace(cer.Name)) <= 0 {
    errStr = errStr + "AccountCertificate.Name is Empty\n"
  }

  if len(strings.TrimSpace(cer.Authority)) <= 0 {
    errStr = errStr + "AccountCertificate.Authority is Empty\n"
  }

  if len(strings.TrimSpace(cer.License)) <= 0 {
    errStr = errStr + "AccountCertificate.License is Empty\n"
  }

  if len(strings.TrimSpace(cer.Url)) <= 0 {
    errStr = errStr + "AccountCertificate.Url is Empty\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (cer AccountCertificate) insertValidation(session *util.Session) error {
  return cer.dataValidation(session)
}

func (cer AccountCertificate) updateValidation(session *util.Session) error {
  err := cer.accountValidation(session)
  if err != nil{
    return err
  }
  return cer.dataValidation(session)
}

func (cer AccountCertificate) deleteValidation(session *util.Session) error {
  return cer.accountValidation(session)
}

func loadAccountCertificates(session *util.Session) ([]AccountCertificate, error) {
  query :=
    "SELECT ID, Name, Authority, License, Url, Description " +
    "FROM AccountCertificate " +
    " WHERE AccountID = $1";

    var result = make([]AccountCertificate, 0)

    certificates, err := util.Select(readAccountCertificate, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummyCer := range certificates {
      certificate, _ := dummyCer.(AccountCertificate)
      result = append(result, certificate)
    }

    return result, nil
}

func readAccountCertificate(rows *sql.Rows) (interface{}, error) {

    var cer AccountCertificate = AccountCertificate{}
    err := rows.Scan(&cer.Id, &cer.Name, &cer.Authority, &cer.License, &cer.Url, &cer.Description)

    return cer, err
}

func (cer AccountCertificate) save(session *util.Session) error {
  if cer.Id <= 0 {
    return cer.saveNew(session)
  }
  return cer.saveUpdate(session)
}

func (cer AccountCertificate) saveNew(session *util.Session) error {

  err := cer.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountCertificate" +
    "(Name, Authority, License, Url, Description, accountID) " +
    "VALUES($1, $2, $3, $4, $5, $6) " +
    "returning ID"

  id, err := util.Insert(query, cer.Name, cer.Authority, cer.License, cer.Url, cer.Description, session.GetAccountID())

  if err != nil {
    return err
  }

  cer.Id = id
  return nil
}

func (cer AccountCertificate) saveUpdate(session *util.Session) error {
  err := cer.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountCertificate " +
    "SET " +
    "Name = $1, " +
    "Authority = $2, " +
    "License = $3, " +
    "Url = $4, " +
    "Description = $5 " +
    "WHERE ID = $6 "

  err = util.Update(query, cer.Name, cer.Authority, cer.License, cer.Url, cer.Description, cer.Id)

  return err

}

func (cer AccountCertificate) delete(session *util.Session) error {
  err := cer.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountCertificate " +
    "WHERE ID = $1 "

  err = util.Update(query, cer.Id)

  return err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveCertificate(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  cerMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Certificate", http.StatusBadRequest}
  }

  cer, _ := parseAccountCertificate(cerMap)

  err = cer.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving Certificate", http.StatusBadRequest}
  }

  return cer, nil
}

func DeleteCertificate(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  cerMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Certificate", http.StatusBadRequest}
  }

  cer, _ := parseAccountCertificate(cerMap)

  err = cer.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting Certificate", http.StatusBadRequest}
  }

  return cer, nil
}

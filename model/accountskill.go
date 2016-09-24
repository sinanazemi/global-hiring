package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type AccountSkill struct {
  Id int `json:"id"`
  SkillID int `json:"skillid"`
  Name string `json:"name"`
  MainServiceID int `json:"mainserviceid"`
  Profeciency SkillProfeciency `json:"profeciency"`
}

func parseAccountSkill(data interface{}) (AccountSkill, error) {
  result := AccountSkill{}

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return result, errors.New("looking for a 'map[string]interface{}' to parse an 'AccountSkill', but not found.\n")
  }

  result.Id = util.ParseInteger(dataMap, "id")
  result.SkillID = util.ParseInteger(dataMap, "skillid")
  result.Name = util.ParseString(dataMap, "name")
  result.MainServiceID = util.ParseInteger(dataMap, "mainserviceid")
  result.Profeciency, _ = parseSkillProfeciency(dataMap["profeciency"])

  return result, nil
}

func parseAccountSkills(data interface{}) []AccountSkill {
  result := make([]AccountSkill, 0)

  sArr, ok := data.([]interface{})
  if (!ok) {
    print("looking for a 'm[]interface{}' to parse an array of 'AccountSkill's, but not found.\n")
    return result
  }

  for _ , s := range sArr {
    skill, err := parseAccountSkill(s)
    if (err == nil) {
      result = append(result, skill)
    }
  }
  return result
}

func (acs *AccountSkill) accountValidation(session *util.Session) error {
  return util.CheckDBAccountValidation(session, "AccountSkill", "AccountID", acs.Id)
}

func (acs *AccountSkill) dataValidation(session *util.Session) error {
  errStr := ""

  if (acs.SkillID <= 0) {
    errStr = errStr + "AccountSkill.SkillID is not valid\n"
  }

  if (acs.Profeciency.isEmpty()) {
    errStr = errStr + "AccountSkill.Profeciency is not valid\n"
  }

  if (len(errStr) > 0) {
    return errors.New(errStr)
  }

  return nil
}

func (acs *AccountSkill) insertValidation(session *util.Session) error {
  return acs.dataValidation(session)
}

func (acs AccountSkill) updateValidation(session *util.Session) error {
  err := acs.accountValidation(session)
  if err != nil{
    return err
  }
  return acs.dataValidation(session)
}

func (acs *AccountSkill) deleteValidation(session *util.Session) error {
  return acs.accountValidation(session)
}

func loadAccountSkills(session *util.Session) ([]AccountSkill, error) {
  query :=
    " SELECT a.ID, a.SkillID, s.Name, s.MainServiceID, a.Profeciency " +
    " FROM accountskill a " +
    " inner join skill s on a.skillID = s.id " +
    " WHERE a.AccountID = $1";

    var result = make([]AccountSkill, 0)

    skills, err := util.Select(readAccountSkill, query, session.GetAccountID())

    if err != nil {
      return result, err
    }

    for _, dummySkill := range skills {
      skill, _ := dummySkill.(AccountSkill)

      skill.Profeciency = loadSkillProfeciency(skill.Profeciency.Value)

      result = append(result, skill)
    }

    return result, nil
}

func readAccountSkill(rows *sql.Rows) (interface{}, error) {

  var skill AccountSkill = AccountSkill{}
  err := rows.Scan(&skill.Id, &skill.SkillID, &skill.Name, &skill.MainServiceID, &skill.Profeciency.Value)

  return skill, err
}

func (acs *AccountSkill) save(session *util.Session) error {
  if acs.Id <= 0 {
    return acs.saveNew(session)
  }
  return acs.saveUpdate(session)
}

func (acs *AccountSkill) saveNew(session *util.Session) error {

  err := acs.insertValidation(session)
  if err != nil {
    return err
  }

  query :=
    "INSERT INTO AccountSkill" +
    "(AccountID, SkillID, profeciency) " +
    "VALUES($1, $2, $3) " +
    "returning ID"

  id, err := util.Insert(query, session.GetAccountID(), acs.SkillID, acs.Profeciency.Value)

  if err != nil {
    return err
  }

  acs.Id = id
  return nil
}

func (acs *AccountSkill) saveUpdate(session *util.Session) error {
  err := acs.updateValidation(session)
  if err != nil {
    return err
  }

  query :=
    "UPDATE AccountSkill " +
    "SET " +
    "SkillID = $1, " +
    "Profeciency = $2 " +
    "WHERE ID = $3 "

  err = util.Update(query, acs.SkillID, acs.Profeciency.Value, acs.Id)

  return err

}

func (acs *AccountSkill) delete(session *util.Session) error {
  err := acs.deleteValidation(session)
  if err != nil {
    return err
  }

  query :=
    "DELETE FROM AccountSkill " +
    "WHERE ID = $1 "

  err = util.Update(query, acs.Id)

  return err
}

func createAccountSkill(skill Skill) AccountSkill {

  return AccountSkill {
    0,
    skill.Id,
    skill.Name,
    skill.MainServiceID,
    skill.Profeciency.clone()  }

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveSkill(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  skillMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Skill", http.StatusBadRequest}
  }

  acs, _ := parseAccountSkill(skillMap)

  err = acs.save(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving skill", http.StatusBadRequest}
  }

  return acs, nil
}

func DeleteSkill(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  skillMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Skill", http.StatusBadRequest}
  }

  acs, _ := parseAccountSkill(skillMap)

  err = acs.delete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while deleting skill", http.StatusBadRequest}
  }

  return acs, nil
}

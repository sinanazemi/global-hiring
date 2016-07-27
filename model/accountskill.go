package model

import(
  "errors"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

const SKILL_PROFECIENCY_STUDENT string = "S"
const SKILL_PROFECIENCY_JUNIOR string = "J"
const SKILL_PROFECIENCY_EXPERIENCED string = "E"
const SKILL_PROFECIENCY_MANAGER string = "M"

type AccountSkill struct {
  Id int `json:"accountskillid"`
  SkillID int `json:"id"`
  Name string `json:"name"`
  MainServiceID int `json:"mainserviceid"`
  IsSelected bool `json:"isselected"`
  Profeciency string `json:"profeciency"`
}

func parseAccountSkill(dataMap map[string]interface{}) (AccountSkill, error) {
  result := AccountSkill{}

  result.Id = util.ParseInteger(dataMap, "accountskillid")
  result.SkillID = util.ParseInteger(dataMap, "id")
  result.Name = util.ParseString(dataMap, "name")
  result.MainServiceID = util.ParseInteger(dataMap, "mainserviceid")
  result.IsSelected = util.ParseBool(dataMap, "isselected")
  result.Profeciency = util.ParseString(dataMap, "profeciency")

  return result, nil
}

func parseAccountSkills(skillsArr []interface{}) []AccountSkill {
  result := make([]AccountSkill, 0)

  for _ , s := range skillsArr {
    smap := s.(map[string]interface{})
    skill, err := parseAccountSkill(smap)
    if err == nil {
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

  profeciencyCheck := false
  profeciencyCheck = profeciencyCheck || (acs.Profeciency == SKILL_PROFECIENCY_STUDENT)
  profeciencyCheck = profeciencyCheck || (acs.Profeciency == SKILL_PROFECIENCY_JUNIOR)
  profeciencyCheck = profeciencyCheck || (acs.Profeciency == SKILL_PROFECIENCY_EXPERIENCED)
  profeciencyCheck = profeciencyCheck || (acs.Profeciency == SKILL_PROFECIENCY_MANAGER)

  if (!profeciencyCheck) {
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
      result = append(result, skill)
    }

    return result, nil
}

func readAccountSkill(rows *sql.Rows) (interface{}, error) {

  var skill AccountSkill = AccountSkill{}
  err := rows.Scan(&skill.Id, &skill.SkillID, &skill.Name, &skill.MainServiceID, &skill.Profeciency)

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

  id, err := util.Insert(query, session.GetAccountID(), acs.SkillID, acs.Profeciency)

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

  err = util.Update(query, acs.SkillID, acs.Profeciency, acs.Id)

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

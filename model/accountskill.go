package model

import(
  "github.com/sinanazemi/global-hiring/util"
)

const SKILL_PROFECIENCY_STUDENT string = "S"
const SKILL_PROFECIENCY_JUNIOR string = "J"
const SKILL_PROFECIENCY_EXPERIENCED string = "E"
const SKILL_PROFECIENCY_MANAGER string = "M"

type AccountSkill struct {
  Id int
  Skill Skill
  Account Account
  Profeciency string
}

func (acs AccountSkill) save() error {
  query :=
    "INSERT INTO AccountSkill" +
    "(AccountID, SkillID, profeciency) " +
    "VALUES($1, $2, $3) " +
    "returning ID"

  id, err := util.Insert(query, acs.Account.Id, acs.Skill.Id, acs.Profeciency)

  if err != nil {
    return err
  }

  acs.Id = id
  return nil
}

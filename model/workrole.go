package model

import(
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type WorkRole struct{
  Value string `json:"value"`
  Name string `json:"name"`
}

var workRoles []WorkRole = make([]WorkRole, 0)

var emptyWorkRole WorkRole = WorkRole{"", ""}

func (w WorkRole) isEmpty() bool {
  return len(w.Value) <= 0
}

func (w WorkRole) clone() WorkRole {
  return WorkRole{w.Value, w.Name}
}

func getEmptyWorkRole() WorkRole {
  return emptyWorkRole.clone()
}

func loadWorkRoles() {
  if len(workRoles) > 0 {
    return
  }

  workRoles = append(workRoles, WorkRole{"I", "Intern"})
  workRoles = append(workRoles, WorkRole{"C", "Individual Contributor"})
  workRoles = append(workRoles, WorkRole{"L", "Lead"})
  workRoles = append(workRoles, WorkRole{"M", "Manager"})
  workRoles = append(workRoles, WorkRole{"E", "Executive"})
  workRoles = append(workRoles, WorkRole{"O", "Owner"})
}

func loadWorkRole(value string) WorkRole {
  loadWorkRoles()

  for _, wr := range workRoles {
    if (value == wr.Value) {
      return wr.clone()
    }
  }
  return getEmptyWorkRole()
}

func parseWorkRole(data interface{}) (WorkRole, error) {

  value, ok := data.(string)
  if (ok) {
    return loadWorkRole(value), nil
  }

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return getEmptyWorkRole(), errors.New("looking for a 'map[string]interface{}' to parse a 'WorkRole', but not found.\n")
  }

  value = util.ParseString(dataMap, "value")
  return loadWorkRole(value), nil

}

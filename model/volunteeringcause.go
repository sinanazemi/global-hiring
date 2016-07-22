package model

import(
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type VolunteeringCause struct{
  Value string `json:"value"`
  Name string `json:"name"`
}

var causes []VolunteeringCause = make([]VolunteeringCause, 0)

var emptyCause VolunteeringCause = VolunteeringCause{"", ""}

func (v VolunteeringCause) isEmpty() bool {
  return len(v.Value) <= 0
}

func (v VolunteeringCause)  clone() VolunteeringCause {
  return VolunteeringCause{v.Value, v.Name}
}

func getEmptyVolunteeringCause() VolunteeringCause {
  return emptyCause.clone()
}

func loadVolunteeringCauses() {
  if len(causes) > 0 {
    return
  }

  causes = append(causes, VolunteeringCause{"Aw", "Animal Welfare"})
  causes = append(causes, VolunteeringCause{"AC", "Arts and Culture"})
  causes = append(causes, VolunteeringCause{"CL", "Children"})
  causes = append(causes, VolunteeringCause{"CS", "Civil Rights and Social Action"})
  causes = append(causes, VolunteeringCause{"DH", "Disaster and Humanitarian Releif"})
  causes = append(causes, VolunteeringCause{"EE", "Economic Empowerment"})
  causes = append(causes, VolunteeringCause{"ED", "Education"})
  causes = append(causes, VolunteeringCause{"EV", "Environment"})
  causes = append(causes, VolunteeringCause{"HL", "Health"})
  causes = append(causes, VolunteeringCause{"HG", "Human Rights"})
  causes = append(causes, VolunteeringCause{"PL", "Politics"})
  causes = append(causes, VolunteeringCause{"PA", "Poverty Alleviation"})
  causes = append(causes, VolunteeringCause{"ST", "Science and technology"})
  causes = append(causes, VolunteeringCause{"SS", "Social Services"})
}

func loadVolunteeringCause(value string) VolunteeringCause {
  loadVolunteeringCauses()

  for _, vc := range causes {
    if (value == vc.Value) {
      return vc.clone()
    }
  }
  return getEmptyVolunteeringCause()
}

func parseVolunteeringCause(data interface{}) (VolunteeringCause, error) {

  value, ok := data.(string)
  if (ok) {
    return loadVolunteeringCause(value), nil
  }

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return getEmptyVolunteeringCause(), errors.New("looking for a 'map[string]interface{}' to parse a 'VolunteeringCause', but not found.\n")
  }

  value = util.ParseString(dataMap, "value")
  return loadVolunteeringCause(value), nil

}

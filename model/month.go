package model

import(
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type Month struct{
  Value int `json:"value"`
  Name string `json:"name"`
}

var months []Month = make([]Month, 0)

var emptyMonth Month = Month{0, ""}

func (m Month) isEmpty() bool {
  return m.Value <= 0
}

func (m Month) clone() Month {
  return Month{m.Value, m.Name}
}

func getEmptyMonth() Month {
  return emptyMonth.clone()
}

func loadMonths() {
  if len(months) > 0 {
    return
  }

  months = append(months, Month{1, "January"})
  months = append(months, Month{2, "February"})
  months = append(months, Month{3, "March"})
  months = append(months, Month{4, "April"})
  months = append(months, Month{5, "May"})
  months = append(months, Month{6, "June"})
  months = append(months, Month{7, "July"})
  months = append(months, Month{8, "August"})
  months = append(months, Month{9, "September"})
  months = append(months, Month{10, "October"})
  months = append(months, Month{11, "November"})
  months = append(months, Month{12, "December"})
}

func loadMonth(value int) Month {
  loadMonths()

  for _, month := range months {
    if (value == month.Value) {
      return month.clone()
    }
  }
  return getEmptyMonth()
}

func parseMonth(data interface{}) (Month, error) {

  valueF, ok := data.(float64)
  if (ok) {
    return loadMonth(int(valueF)), nil
  }

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return getEmptyMonth(), errors.New("looking for a 'map[string]interface{}' to parse a 'Month', but not found.\n")
  }

  value := util.ParseInteger(dataMap, "value")
  return loadMonth(value), nil

}

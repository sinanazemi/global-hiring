package model

import(
  "errors"
  "github.com/sinanazemi/global-hiring/util"
)

type LanguageProfeciency struct{
  Value string `json:"value"`
  Name string `json:"name"`
}

var languageProfeciencies []LanguageProfeciency = make([]LanguageProfeciency, 0)

var emptyLanguageProfeciency LanguageProfeciency = LanguageProfeciency{"", ""}

func (l LanguageProfeciency) isEmpty() bool {
  return len(l.Value) <= 0
}

func (l LanguageProfeciency) clone() LanguageProfeciency {
  return LanguageProfeciency{l.Value, l.Name}
}

func getEmptyLanguageProfeciency() LanguageProfeciency {
  return emptyLanguageProfeciency.clone()
}

func loadLanguageProfeciencies() {
  if len(languageProfeciencies) > 0 {
    return
  }

  languageProfeciencies = append(languageProfeciencies, LanguageProfeciency{"E", "Elementary"})
  languageProfeciencies = append(languageProfeciencies, LanguageProfeciency{"B", "Basic"})
  languageProfeciencies = append(languageProfeciencies, LanguageProfeciency{"C", "Conversational"})
  languageProfeciencies = append(languageProfeciencies, LanguageProfeciency{"F", "Fluent"})
  languageProfeciencies = append(languageProfeciencies, LanguageProfeciency{"N", "Native or Bilingual"})
}

func loadLanguageProfeciency(value string) LanguageProfeciency {
  loadLanguageProfeciencies()

  for _, lp := range languageProfeciencies {
    if (value == lp.Value) {
      return lp.clone()
    }
  }
  return getEmptyLanguageProfeciency()
}

func parseLanguageProfeciency(data interface{}) (LanguageProfeciency, error) {

  value, ok := data.(string)
  if (ok) {
    return loadLanguageProfeciency(value), nil
  }

  dataMap, ok := data.(map[string]interface{})
  if (!ok) {
    return getEmptyLanguageProfeciency(), errors.New("looking for a 'map[string]interface{}' to parse a 'LanguageProfeciency', but not found.\n")
  }

  value = util.ParseString(dataMap, "value")
  return loadLanguageProfeciency(value), nil

}

package mapper

type Education struct {
  School string `json="school"`
  FromDate int `json="fromDate"`
  ToDate int `json="toDate"`
  Degree string `json="degree"`
  Field string `json="field"`
  Grade float32 `json="grade"`
}

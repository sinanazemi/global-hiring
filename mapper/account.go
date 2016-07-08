package mapper

type Account struct {
  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City string `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`

  IsStudent bool `json:isStudent`

  Languages []Language `json="languages"`

  Educations []Education `json="education"`

  MainServices []MainService `json= mainServices`

}

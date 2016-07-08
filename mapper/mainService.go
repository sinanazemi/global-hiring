package mapper

type MainService struct{
  Name string `json="name"`
  Skills []Skill `json="skills"`
}

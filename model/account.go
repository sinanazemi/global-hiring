package model

import (
  "errors"
  "strings"
  "database/sql"
  "net/http"
  "github.com/sinanazemi/global-hiring/util"
)

type Account struct {

  Id     int    `json:"id"`
	Name  string `json:"name"`
  Email string `json:"email"`
  City City `json:"city"`
  Phone string `json:"phone"`
  Password string `json:"password"`
  Description string `json:"description"`
  JobTitle string `json:"jobtitle"`
  IsStudent bool `json:"isstudent"`
  IsCompleted bool `json:"iscompleted"`

  Languages []AccountLanguage `json:"languages"`

  Educations []AccountEducation `json:"educations"`

  Skills []AccountSkill `json:"skills"`

  Certificates []AccountCertificate `json:"certificates"`

  Works []AccountWork `json:"works"`

  Volunteerings []AccountVolunteering `json:"volunteerings"`

  Courses []AccountCourse `json:"courses"`

  Honors []AccountHonor `json:"honors"`

  Tests []AccountTest `json:"tests"`

  Projects []AccountProject `json:"projects"`
}

func parseAccount(dataMap map[string]interface{}) (Account, error) {

  id := 0
  if dataMap["id"] != nil {
    id = int(dataMap["id"].(float64))
  }

  result := Account{Id: id}

  result.Name = util.ParseString(dataMap, "name")
  result.Email = util.ParseString(dataMap, "email")
  result.Phone = util.ParseString(dataMap, "phone")
  result.Password = util.ParseString(dataMap, "password")
  result.Description = util.ParseString(dataMap, "description")
  result.JobTitle = util.ParseString(dataMap, "jobtitle")
  result.IsStudent = util.ParseBool(dataMap, "isstudent")
  result.IsCompleted = util.ParseBool(dataMap, "iscompleted")
  result.City, _ = parseCity(dataMap["city"])

  result.Languages      = parseAccountLanguages(      dataMap["languages"])
  result.Educations     = parseAccountEducations(     dataMap["educations"])
  result.Skills         = parseAccountSkills(         dataMap["skills"])

  result.Certificates   = parseAccountCertificates(   dataMap["certificates"])
  result.Works          = parseAccountWorks(          dataMap["works"])
  result.Volunteerings  = parseAccountVolunteerings(  dataMap["volunteerings"])
  result.Courses        = parseAccountCourses(        dataMap["courses"])
  result.Honors         = parseAccountHonors(         dataMap["honors"])
  result.Tests          = parseAccountTests(          dataMap["tests"])
  result.Projects       = parseAccountProjects(       dataMap["projects"])

  return result, nil
}

func loadAccount(session *util.Session) (Account, error) {
  query := "select ID, Name, Email, Phone, Password, Description, JobTitle, IsStudent, cityID from Account Where ID = $1"
  accArr, err := util.Select(readAccount, query, session.GetAccountID())
  if err != nil {
    return Account{}, err
  }

  account := accArr[0].(Account)

  account.City = loadCity(account.City.Id)

  account.Languages, _ = loadAccountLanguages(session)
  account.Educations, _ = loadAccountEducations(session)
  account.Skills, _ = loadAccountSkills(session)
  account.Certificates, _ = loadAccountCertificates(session)
  account.Works, _ = loadAccountWorks(session)
  account.Volunteerings, _ = loadAccountVolunteerings(session)
  account.Courses, _ = loadAccountCourses(session)
  account.Honors, _ = loadAccountHonors(session)
  account.Tests, _ = loadAccountTests(session)
  account.Projects, _ = loadAccountProjects(session)

  account.IsCompleted = len(account.Languages) > 0

  return account, nil
}

func readAccount(rows *sql.Rows) (interface{}, error) {
  var acc Account = Account{}

  err := rows.Scan(
      &acc.Id, &acc.Name, &acc.Email, &acc.Phone, &acc.Password,
      &util.NullableString{&acc.Description}, &util.NullableString{&acc.JobTitle},
      &acc.IsStudent, &acc.City.Id)
  return acc, err
}

func (acc *Account) initializeJobTitle() {
  acc.JobTitle = acc.Skills[0].Name
}

func (acc *Account) create(session *util.Session) error {

  acc.Email = strings.ToLower(strings.TrimSpace(acc.Email))
  acc.Password = util.GetMD5Hash(acc.Password)

  acc.initializeJobTitle()

  query :=
    "INSERT INTO Account" +
    "(Name, Email, cityID, Phone, Password, Description, JobTilte, isStudent) " +
    "VALUES($1, $2, $3, $4, $5, $6, $7, $8) " +
    "returning ID"

  id, err := util.Insert(query, acc.Name, acc.Email, acc.City.Id, acc.Phone, acc.Password, acc.Description, acc.JobTitle, acc.IsStudent)

  if err != nil {
    return err
  }
  acc.Id = id
  session.PutAccountID(id)

  return acc.createComplete(session)
}

func (acc *Account) createComplete(session *util.Session) error {

    if session.GetAccountID() != acc.Id {
      return errors.New("It seems that this account is not authenticated in this session")
    }

    for _ , language := range acc.Languages {
      language.save(session)
    }

    for _ , education := range acc.Educations {
      education.save(session)
    }

    for _ , skill := range acc.Skills {
      skill.save(session)
    }

    for _ , certificate := range acc.Certificates {
      certificate.save(session)
    }

    for _ , work := range acc.Works {
      work.save(session)
    }

    for _ , vol := range acc.Volunteerings {
      vol.save(session)
    }

    for _ , course := range acc.Courses {
      course.save(session)
    }

    for _ , honor := range acc.Honors {
      honor.save(session)
    }

    for _ , test := range acc.Tests {
      test.save(session)
    }

    for _ , project := range acc.Projects {
      project.save(session)
    }

    return nil
}

func (acc *Account) updateDescription (session *util.Session, dataMap map[string]interface{}) error {

  if session.GetAccountID() != acc.Id {
    return errors.New("It seems that this account is not authenticated in this session")
  }

  err := util.CheckString(dataMap, "description")
  if err != nil {
    return err
  }

  description := util.ParseString(dataMap, "description")

  query :=
  "UPDATE Account " +
  "SET Description = $1 " +
  "WHERE ID = $2 "

  err = util.Update(query, description, acc.Id)

  if err != nil {
    return err
  }

  acc.Description = description

  return nil
}

func (acc *Account) updateJobTitle (session *util.Session, dataMap map[string]interface{}) error {

  if session.GetAccountID() != acc.Id {
    return errors.New("It seems that this account is not authenticated in this session")
  }

  err := util.CheckString(dataMap, "jobtitle")
  if err != nil {
    return err
  }

  jobtitle := util.ParseString(dataMap, "jobtitle")

  query :=
  "UPDATE Account " +
  "SET JobTitle = $1 " +
  "WHERE ID = $2 "

  err = util.Update(query, jobtitle, acc.Id)

  if err != nil {
    return err
  }

  acc.JobTitle = jobtitle

  return nil
}

func (acc *Account) getStrength() int {
  var result int = 0

  //adding profile summary, each character +0.2
  //result = result + ?

  result = result + getWorkStrength(acc.Works)
  result = result + getLanguageStrength(acc.Languages)
  result = result + getEducationStrength(acc.Educations)
  result = result + getCertificateStrength(acc.Certificates)
  result = result + getHonorStrength(acc.Honors)
  result = result + getProjectStrength(acc.Projects)
  result = result + getTestStrength(acc.Tests)
  result = result + getVolunteeringStrength(acc.Volunteerings)
  result = result + getCourseStrength(acc.Courses)

  return result
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  account, err := loadAccount(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while loading account", http.StatusBadRequest}
  }
  account.Password = "" // :D
  return account, nil
}

func SaveAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  accountMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, _ := parseAccount(accountMap)

  err = account.create(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account", http.StatusBadRequest}
  }

  return GetAccount(w, r)
}

func CompleteAccount(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  accountMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, _ := parseAccount(accountMap)

  err = account.createComplete(session)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account", http.StatusBadRequest}
  }

  return GetAccount(w, r)
}

func GetAccountStrength(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return -1, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  account, err := loadAccount(session)
  if err != nil {
    return -1, &util.HandlerError{err, "Problem while loading account", http.StatusBadRequest}
  }
  return account.getStrength(), nil
}

func SaveDescription(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  descMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, err := loadAccount(session)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid Account in this session", http.StatusBadRequest}
  }

  err = account.updateDescription(session, descMap)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account Description", http.StatusBadRequest}
  }

  return account, nil
}

func SaveJobTitle(w http.ResponseWriter, r *http.Request) (interface{}, *util.HandlerError) {

  session, err := util.GetSession(w, r)
  if err != nil {
      return nil, &util.HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  jobMap, err := util.ParseJsonRequest(r)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  account, err := loadAccount(session)
  if err != nil {
      return nil, &util.HandlerError{err, "Invalid Account in this session", http.StatusBadRequest}
  }

  err = account.updateJobTitle(session, jobMap)
  if err != nil {
    return nil, &util.HandlerError{err, "Problem while saving account Job Title", http.StatusBadRequest}
  }

  return account, nil
}

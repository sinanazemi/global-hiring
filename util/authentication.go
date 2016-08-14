package util

import (
  "errors"
  "strings"
  "net/http"
)

type Redirect struct {
	Url string `json:"url"`
	DisplayOK bool `json:"displayok"`
}

func RedirectCheck(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {

	red := Redirect{}

	session, err := GetSession(w, r)
  if err != nil {
      return nil, &HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  m, err := ParseJsonRequest(r)
  if err != nil {
      return nil, &HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

	url := m["url"].(string)

  login := (session.GetAccountID() > 0)

	if (url == "login.html") {
		red.DisplayOK= !login
		if login {
			red.Url="profile.html"
		}
	}

	if (url == "profile.html") {
		red.DisplayOK= login
		if !login {
			print("\nlogin.html\n")
			red.Url="login.html"
		}
	}

	return red, nil
}

func Authenticate(w http.ResponseWriter, r *http.Request) (int, *HandlerError) {

  id := -1

	session, err := GetSession(w, r)
  if err != nil {
      return id, &HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  m, err := ParseJsonRequest(r)
  if err != nil {
      return id, &HandlerError{err, "Invalid JSON Account", http.StatusBadRequest}
  }

  email, ok := m["email"].(string)
  if !ok {
    msg := "Looking for a 'email' to authenticate!"
    return id, &HandlerError{errors.New(msg), msg, http.StatusBadRequest}
  }
  email = strings.ToLower(strings.TrimSpace(email))

  password, ok := m["password"].(string)
  if !ok {
    msg := "Looking for a 'password' to authenticate!"
    return id, &HandlerError{errors.New(msg), msg, http.StatusBadRequest}
  }
  password = GetMD5Hash(password)

  print(email + " - " + password + "\n")

  id , _ = SelectInteger("select id from Account where Email = $1 and password = $2", email, password)

  if (id <= 0) {
    msg := "Invalid email or password!"
    return id, &HandlerError{errors.New(msg), msg, http.StatusBadRequest}
  }

  session.PutAccountID(id)

  return id, nil
}

type DummyStruct struct{

}

func Logout(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {

	session, err := GetSession(w, r)
  if err != nil {
      return DummyStruct{}, &HandlerError{err, "Problems in session", http.StatusBadRequest}
  }

  session.clearAccountID()

  return DummyStruct{}, nil
}

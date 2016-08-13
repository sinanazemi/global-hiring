package util

import (
    "net/http"
    "encoding/base64"
    "github.com/gorilla/sessions"
)

const (
  sessionName string = "session-name"
  accountKey string = "accountKeyForGlobalHiring"
)

var (
  cookieStoreKey, _ = base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
  store = sessions.NewCookieStore(cookieStoreKey)
)

type Session struct {
  request *http.Request
  response http.ResponseWriter
  session *sessions.Session
}

func GetSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
    // Get a session. We're ignoring the error resulted from decoding an
    // existing session: Get() always returns a session, even if empty.
    session, err := store.Get(r, sessionName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return nil, err
    }

    session.Options = &sessions.Options{
    Path:     "/",
    MaxAge:   60 * 10, // 10 minutes
    HttpOnly: true,
}

    return &Session{r, w, session}, nil
}

func (session Session) Get (key interface{}) interface{} {
  return session.session.Values[key]
}

func (session Session) Put (key interface{}, value interface{}) {

  session.session.Values[key] = value
  session.session.Save(session.request, session.response)
}

func (session Session) GetAccountID() int {
  idVal := session.Get(accountKey)
  id, ok := idVal.(int)
  if !ok {
    return -1
  }
  return id
}

func (session Session) PutAccountID(accountID int) {
  session.Put(accountKey, accountID)
}

func (session Session) clearAccountID() {
  session.Put(accountKey, nil)
}

func (session Session) Refresh() {
  id := session.GetAccountID()
  if (id > 0){
    session.PutAccountID(id)
  }
}

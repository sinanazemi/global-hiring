package util

import (
    "database/sql"
    "fmt"
    "errors"
    _ "github.com/lib/pq"
    //"time"
)

const (
    DB_USER     = "globeadmin"
    DB_PASSWORD = "globePassword"
    DB_NAME     = "globalhiring"
)

var dbinfo string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
    DB_USER, DB_PASSWORD, DB_NAME)

var db *sql.DB = nil

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func checkConnection() error {
  if db == nil {
    return initConnection()
  }

  err := db.Ping()
  if err == nil {
    return nil
  }

  err = db.Close()
  checkErr(err)

  return initConnection()
}

func initConnection() error {

  var err error = nil

  db, err = sql.Open("postgres", dbinfo)
  checkErr(err)

  return err
}

type MapperFunc func(*sql.Rows) (interface{}, error)

func Select(mapper MapperFunc, query string, args ...interface{}) ([]interface{}, error) {

  err := checkConnection()
  checkErr(err)
  if err != nil {
    return nil, err
  }

  rows, err := db.Query(query, args...)
  checkErr(err)
  if err != nil {
    return nil, err
  }

  var result = make([]interface{}, 0)

  for rows.Next() {
      data, err := mapper(rows)
      if err != nil {
        return nil, err
      }
      result = append(result, data)
  }
  return result, nil
}

func SelectInteger(query string, args ...interface{}) (int, error) {

  result := -1

  err := checkConnection()
  checkErr(err)
  if err != nil {
    return result, err
  }

  rows, err := db.Query(query, args...)
  checkErr(err)
  if err != nil {
    return result, err
  }

  for rows.Next() {
    err = rows.Scan(&result)
  }
  return result, err
}

func Insert(query string, args ...interface{}) (int, error) {

  err := checkConnection()
  checkErr(err)
  if err != nil {
    return -1, err
  }

  var InsertedId int
  // "INSERT INTO MyTable(MyTableID, name) VALUES($1,$2) returning MyTableID;"
  err = db.QueryRow(query, args...).Scan(&InsertedId)
  checkErr(err)
  if err != nil {
    return -1, err
  }

  return InsertedId, nil
}

func Update(query string, args ...interface{}) error {

  err := checkConnection()
  checkErr(err)
  if err != nil {
    return err
  }

  stmt, err := db.Prepare(query)
  checkErr(err)
  if err != nil {
    return err
  }

  res, err := stmt.Exec(args...)
  checkErr(err)
  if err != nil {
    return err
  }

  affect, err := res.RowsAffected()
  checkErr(err)
  if err != nil {
    return err
  }

  if (affect <= 0) {
    return errors.New("No rows updated!")
  }

  return nil
}

/*func main() {

    var lastInsertId int
    err = db.QueryRow("INSERT INTO MyTable(MyTableID, name) VALUES($1,$2) returning MyTableID;", 20, "astaxie").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    fmt.Println("# Updating")
    stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created time.Time
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("uid | username | department | created ")
        fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
    }

    fmt.Println("# Deleting")
    stmt, err = db.Prepare("delete from userinfo where uid=$1")
    checkErr(err)

    res, err = stmt.Exec(lastInsertId)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")
}*/


func CheckDBAccountValidation(session *Session, tableName string, accountColumnName string, recordID int) error {

  err := checkConnection()
  checkErr(err)
  if err != nil {
    return err
  }

  query := "select count(*) from " + tableName + " where " + accountColumnName + " = $1 and id = $2";

  rows, err := db.Query(query, session.GetAccountID(), recordID)
  checkErr(err)

  for rows.Next() {
    var count int
    err = rows.Scan(&count)
    checkErr(err)
    if (count > 0) {
      return nil
    }
  }
  return errors.New("Authorization Faild for updating/removing a record in AccountSkill")
}

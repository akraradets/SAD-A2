package machine

import (
  "fmt"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
  "log"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "vendingApp"
  password = "vendingPassword"
  dbname   = "vendingDB"
)

var DB *sqlx.DB

func InitDb() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  var err error
  DB, err = sqlx.Connect("postgres", psqlInfo)
  if err != nil {
      log.Panic(err)
  }

  if err = DB.Ping(); err != nil {
      log.Panic(err)
  }
}

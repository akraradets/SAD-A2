package main

import (
  "database/sql"
  "fmt"
  "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "vendingApp"
  password = ""
  dbname   = "vendingDB"
)

func test_connection() (string, error) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()


  return "Successfully connected!", err
}

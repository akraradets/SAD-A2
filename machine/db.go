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

//variable to cache items in the list : cache aside pattern
var itemsCache map[string]Item

func LoadItems() {
  items := []Item{}
  err := DB.Select(&items, "SELECT * FROM items")
  if err != nil {
    log.Panic(err)
  }
  itemsCache = make(map[string]Item)
  for _, item := range items {
    itemsCache[item.Name] = item
  }

  log.Println(itemsCache["Coca-cola"].Name)
  log.Println(itemsCache["Pringles"].Name)
}

func UpdateItems() {
  for _, item := range itemsCache {
    _, err := DB.NamedExec(`UPDATE items SET amount=:amount WHERE name=:name`, item)
    if err != nil {
      log.Panic(err)
    }
  }
}

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

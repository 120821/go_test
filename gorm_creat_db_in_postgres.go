package main

import (
	"gorm.io/driver/postgres"
   "log"
   _ "github.com/lib/pq"
)

func main() {
  conninfo := "user=admin password=88888888 host=127.0.0.1 sslmode=disable"
  db, err := sql.Open("postgres", conninfo)

  if err != nil {
    log.Fatal(err)
  }
  dbName := "testdb"
  _, err = db.Exec("create database " + dbName)
  if err != nil {
    //handle the error
    log.Fatal(err)
  }

  //Then execute your query for creating table
  _, err = db.Exec("CREATE TABLE example ( id integer,
  username varchar(255) )")

  if err != nil {
    log.Fatal(err)
  }

}

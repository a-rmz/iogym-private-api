package db

import (
  "fmt"
  "os"
  "database/sql"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)


func Connect() *sqlx.DB {
  host     := os.Getenv("DB_HOST")
  port     := os.Getenv("DB_PORT")
  user     := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  dbname   := os.Getenv("DB_NAME")

  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sqlx.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  return db
}

func QueryMany(query string, models []interface{}, params interface{}) {
  db := Connect()
  defer db.Close()

  rows, _ := db.NamedQuery(query, params)
  i := 0

  for rows.Next() {
    switch err := rows.StructScan(&models[i]); err {
    case sql.ErrNoRows:
      fmt.Println("No rows were returned!")
    case nil:
      i += 0
      fmt.Println(models[i])
    default:
      panic(err)
    }
  }
}

func QueryOne(query string, params map[string]interface{}) *sqlx.Rows {
  db := Connect()
  defer db.Close()

  rows, _ := db.NamedQuery(query, params)
  
  if rows != nil {
    return rows
  }
  return nil
}

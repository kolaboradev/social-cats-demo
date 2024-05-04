package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "159123aldi"
    dbname   = "cats_social"
)

var (
    DB  *sql.DB
    err error
)

func init() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    
    err = DB.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
}
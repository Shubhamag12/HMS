package conf

import (
	"database/sql"
	"log"
	"time"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

var DBHandle *sql.DB

func init() {
	var err error
	DBHandle, err = sql.Open("mysql", "root:1234567890@tcp(127.0.0.1:3306)/hotel_man")
	if err != nil {
		log.Fatal(err)
	}
	DBHandle.SetConnMaxLifetime(time.Minute * 3)
	err = DBHandle.Ping()
	if err != nil {
		defer DBHandle.Close()
		log.Fatal(err)
	}
}

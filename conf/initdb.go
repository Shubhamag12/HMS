package conf

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBHandle *sql.DB

func init() {
	connString := fmt.Sprintf("%v:%v@%v(%v:%v)/%v", EnvMap["DB_USER"], EnvMap["DB_PASSWORD"], EnvMap["DB_PROTO"], EnvMap["DB_HOST"], EnvMap["DB_PORT"], EnvMap["DB_NAME"])
	var err error
	DBHandle, err = sql.Open("mysql", connString)
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

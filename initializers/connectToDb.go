package initializers

import (
	"database/sql"
	"os"
)

var DB *sql.DB

func ConnectToDB() {
	var err error
	dbDriver := os.Getenv("dbDriver")
	dbSource := os.Getenv("dbSource")
	DB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		panic("Faild to connect to db")
	}
}

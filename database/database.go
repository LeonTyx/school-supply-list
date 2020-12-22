package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	DBCon        *sql.DB
)

func InitDB()  {
	var err error

	// Open up our database connection.
	DBCon, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Cannot open SQL connection")
		panic(err.Error())
	}
	fmt.Println("Successful database connection!")
}
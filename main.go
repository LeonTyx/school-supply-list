package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"school-supply-list/api"
	"school-supply-list/database"
)

//Load the enviroment variables from the projectvars.env file
func initEnv()  {
	err := godotenv.Load("projectvars.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//Check that database is up to date.
//Will cycle through all changes in db/migrations until the database is up to date
func PerformMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	fmt.Println("Database migrations completed. Database should be up to date")
}

//Initialize a database connection using the environment variable DATABASE_URL
//Returns type *sql.DB
func initDBConnection() *sql.DB{
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Cannot open SQL connection")
		panic(err.Error())
	}

	return db
}

func main() {
	initEnv()
	PerformMigrations()
	db := initDBConnection()

	r := gin.Default()
	dbConnection := &database.DB{Db: db}
	r.GET("/ping", api.Test(dbConnection))
	_ = r.Run()
}
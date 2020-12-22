package main

import (
	"database/sql"
	"fmt"
	_ "github.com/antonlindstrom/pgstore"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"school-supply-list/api"
	"school-supply-list/database"
)


func initEnv()  {
	err := godotenv.Load("projectvars.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Cannot open SQL connection")
		panic(err.Error())
	}

	r := gin.Default()

	dbConnection := &database.DB{Db: db}
	r.GET("/ping", api.Test(dbConnection))
	_ = r.Run()
}
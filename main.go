package main

import (
	"./database"
	"github.com/joho/godotenv"
	"log"
)

func initEnv()  {
	err := godotenv.Load("projectvars.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()
	database.InitDB()
}

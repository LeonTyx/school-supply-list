package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"school-supply-list/api"
	"school-supply-list/auth/authentication"
	"school-supply-list/database"
	"time"
)

//Load the environment variables from the projectvars.env file
func initEnv() {
	if _, err := os.Stat("projectvars.env"); err == nil {
		err := godotenv.Load("variables.env")
		if err != nil {
			fmt.Println("Error loading environment.env")
		}
		fmt.Println("Current environment:", os.Getenv("ENV"))
	}
}

func createServer(dbConnection *database.DB) *gin.Engine {
	r := gin.Default()

	authentication.Routes(r.Group("oauth/v1"), dbConnection)
	api.Routes(r.Group("api/v1"), dbConnection)

	return r
}

func main() {
	initEnv()
	database.PerformMigrations()
	authentication.ConfigOauth()
	db := database.InitDBConnection()
	defer db.Close()

	SStore := database.InitOauthStore()
	// Run a background goroutine to clean up expired sessions from the database.
	defer SStore.StopCleanup(SStore.Cleanup(time.Minute * 5))
	dbConnection := &database.DB{Db: db, SessionStore: SStore}

	r := createServer(dbConnection)

	_ = r.Run()
}

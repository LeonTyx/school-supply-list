package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"school-supply-list/api"
	"school-supply-list/auth/authentication"
	"school-supply-list/database"
	"time"
)

//Load the environment variables from the projectvars.env file
func initEnv() {
	err := godotenv.Load("projectvars.env")
	if err != nil {
		log.Fatal(err)
	}
}

func createServer(dbConnection *database.DB) *gin.Engine{
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

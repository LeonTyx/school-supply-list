package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
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

//Check that database is up to date.
//Will cycle through all changes in db/migrations until the database is up to date
func PerformMigrations() {
	m, err := migrate.New(
		"file://database/migrations",
		os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	fmt.Println("Database migrations completed. Database should be up to date")
}

func main() {
	initEnv()
	PerformMigrations()
	authentication.ConfigOauth()
	db := database.InitDBConnection()
	defer db.Close()
	SStore := database.InitOauthStore()

	// Run a background goroutine to clean up expired sessions from the database.
	defer SStore.StopCleanup(SStore.Cleanup(time.Minute * 5))
	r := gin.Default()
	dbConnection := &database.DB{Db: db, SessionStore: SStore}

	authentication.Routes(r.Group("oauth/v1"), dbConnection)
	api.Routes(r.Group("api/v1"), dbConnection)

	_ = r.Run()
}

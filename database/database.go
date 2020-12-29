package database

import (
	"database/sql"
	"fmt"
	"github.com/antonlindstrom/pgstore"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"net/http"
	"os"
)

type DB struct {
	Db           *sql.DB
	SessionStore *pgstore.PGStore
}

//Initialize a database connection using the environment variable DATABASE_URL
//Returns type *sql.DB
func InitDBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Cannot open SQL connection")
		panic(err.Error())
	}

	return db
}

func InitOauthStore() *pgstore.PGStore {
	var err error

	SessionStore, err := pgstore.NewPGStore(os.Getenv("DATABASE_URL"), []byte(os.Getenv("DATABASE_SECRET")))
	if err != nil {
		panic(err)
	}

	SessionStore.MaxAge(1800)
	SessionStore.Options.SameSite = http.SameSiteLaxMode
	SessionStore.Options.HttpOnly = true
	if os.Getenv("ENV") == "DEV" {
		SessionStore.Options.Secure = false
	} else {
		SessionStore.Options.Secure = true
	}
	fmt.Println("Successful oauth store connection!")
	return SessionStore
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
package database

import "database/sql"

type DB struct {
	Db   *sql.DB
}
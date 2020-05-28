package database

import (
	"database/sql"
	"log"

	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"

	//
	_ "github.com/golang-migrate/migrate/source/file"
)

// DB is a pointer to the database handler
var DB *sql.DB

// InitDB creates a connection to our database.
func InitDB() {
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

// Migrate runs migrations file for us.
func Migrate() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

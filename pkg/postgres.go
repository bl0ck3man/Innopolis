package postgres

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	db         *sqlx.DB
	Connection = newConfig()
)

const (
	MaxOpenConns = 25
	MaxIdleConns = 60 * int(time.Second)
)

func newConfig() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occurred by reading env file")
	}

	if db != nil {
		return db
	}

	db, err := sqlx.Connect("pgx", os.Getenv("PG_CONNECT"))
	if err != nil {
		log.Fatalln(err)
	}
	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	return db
}

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:a123456@localhost:5432/gen_laravel?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

var err error

func TestMain(m *testing.M) {
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connect sql fail ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}

package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	devDriver = "postgres"
	devSource = "postgresql://root:a123456@localhost:5432/gen_laravel?sslmode=disable"
)

func ConnDev() *Queries {
	testDB, err := sql.Open(devDriver, devSource)
	//defer testDB.Close()
	if err != nil {
		log.Fatal("connect sql fail ", err)
	}
	DQueries := New(testDB)
	return DQueries
}

func TrunateDB() {
	testDB, err := sql.Open(devDriver, devSource)
	defer testDB.Close()
	if err != nil {
		log.Fatal("connect sql fail ", err)
	}
	Queries := New(testDB)
	fmt.Println("trance table")
	Queries.TruncateTbField(context.Background())
	Queries.TruncateTB(context.Background())
	Queries.TruncateProject(context.Background())
}

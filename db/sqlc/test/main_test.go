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
	dbSource = "postgresql://root:root@123@localhost:5432/simpleBank?sslmode=disable"
)

var testQueries *Queries

func Testmain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connection to db err:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}

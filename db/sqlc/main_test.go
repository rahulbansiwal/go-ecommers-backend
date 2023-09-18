package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testdb *sql.DB

func TestMain(m *testing.M) {
	//TODO: load config data from app.env using viper
	testdb, err := sql.Open("postgres", "postgresql://rahul:admin@localhost:5432/ecom?sslmode=disable")
	if err != nil {
		log.Fatal("connection to DB can't be established", err)
	}
	testQueries = New(testdb)
	os.Exit(m.Run())
}

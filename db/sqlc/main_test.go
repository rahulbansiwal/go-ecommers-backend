package sqlc

import (
	"database/sql"
	"ecom/db/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testdb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't load config values from the file", err)
	}
	testdb, err = sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("connection to DB can't be established", err)
	}
	testQueries = New(testdb)
	os.Exit(m.Run())
}

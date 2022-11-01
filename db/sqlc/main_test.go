package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	util "github.com/dhiyaaulauliyaa/learn-go/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	/* Load Config File */
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error when connecting to databse: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

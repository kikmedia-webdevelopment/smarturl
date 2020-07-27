package storage

import (
	"os"
	"testing"

	"github.com/juliankoehn/mchurl/config"
)

var (
	testDbName = "database.db"
)

func TestDialSqlite(t *testing.T) {
	sqliteConfig := &config.Configuration{
		DB: config.DBConfiguration{
			Driver: "sqlite",
			URL:    testDbName,
		},
	}

	_, err := Dial(sqliteConfig)
	if err != nil {
		t.Error(err)
	}

	// clean up db file from disk after test
	os.Remove(testDbName)
}

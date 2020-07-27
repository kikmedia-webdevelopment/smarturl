package models

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AuditTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (a *AuditTestSuite) TestAudit() {
	err := NewAuditLogEntry(a.db, &User{
		ID:    1,
		Email: "me+testing@julian.pro",
	}, LoginAction)

	assert.Equal(a.T(), err, nil)
}

func TestAuditTestSuite(t *testing.T) {
	config := &config.Configuration{
		DB: config.DBConfiguration{
			Driver: "sqlite",
			URL:    testDbName,
		},
	}
	db, err := storage.Dial(config)
	require.NoError(t, err)

	defer db.Close()
	db.AutoMigrate(&AuditLogEntry{})

	ts := &AuditTestSuite{
		db: db,
	}
	suite.Run(t, ts)

	// clean up db file from disk after test
	os.Remove(testDbName)
}

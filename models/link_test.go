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

type LinkTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (l *LinkTestSuite) TestCreateLink() {
	link, err := CreateEntry(l.db, 4, "https://www.test.de", "")
	assert.Equal(l.T(), err, nil)
	assert.NotEmpty(l.T(), link.ID)

	LinkWithID, err := CreateEntry(l.db, 4, "https://www.google.de", "GOOGL")
	assert.Equal(l.T(), err, nil)
	assert.Equal(l.T(), LinkWithID.ID, "GOOGL")
}

func (l *LinkTestSuite) TestGetLinkByID() {
	link, err := GetLinkByID(l.db, "GOOGL")
	assert.Equal(l.T(), err, nil)
	assert.Equal(l.T(), link.ID, "GOOGL")

	links, err := LinksList(l.db)
	assert.Equal(l.T(), err, nil)
	assert.Equal(l.T(), len(links), 2) // we created 2 in TestCreateLink()

	link.URL = "https://www.julian.pro"
	newLink, err := LinkUpdate(l.db, link)
	assert.Equal(l.T(), err, nil)
	assert.Equal(l.T(), newLink.URL, "https://www.julian.pro")

	// increase counter
	err = IncreaseVisitCounter(l.db, link)
	assert.Equal(l.T(), err, nil)
	// found, lets delete
	err = DeleteEntry(l.db, "GOOGL")
	assert.Equal(l.T(), err, nil)
}

func TestLinkTestSuite(t *testing.T) {
	config := &config.Configuration{
		DB: config.DBConfiguration{
			Driver: "sqlite",
			URL:    testDbName,
		},
	}
	db, err := storage.Dial(config)
	require.NoError(t, err)

	defer db.Close()
	db.AutoMigrate(&Link{})

	ts := &LinkTestSuite{
		db: db,
	}
	suite.Run(t, ts)

	// clean up db file from disk after test
	os.Remove(testDbName)
}

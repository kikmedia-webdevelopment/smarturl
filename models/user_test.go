package models

import (
	"os"
	"testing"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	testDbName = "database.db"
)

type UserTestSuite struct {
	suite.Suite
	db *storage.Connection
}

func (u *UserTestSuite) TestUser() {
	user := &User{
		Email:    "me+testing@julian.pro",
		Password: "testPassword",
	}

	err := u.db.Create(user).Error
	assert.Equal(u.T(), err, nil)

	err = user.SetEmail(u.db, "me+testing2@julian.pro")
	assert.Equal(u.T(), err, nil)

	err = user.UpdatePassword(u.db, "newPassword")
	assert.Equal(u.T(), err, nil)

	ok := user.Authenticate("newPassword")
	assert.Equal(u.T(), ok, true)

	// test ConfirmEmailChange
	// test Recover

	duplicate, err := IsDuplicatedEmail(u.db, "me+testing2@julian.pro")
	assert.Equal(u.T(), err, nil)
	assert.Equal(u.T(), duplicate, true)

	noDupli, err := IsDuplicatedEmail(u.db, "me+testing@julian.pro")
	assert.Equal(u.T(), err, nil)
	assert.Equal(u.T(), noDupli, false)

	foundUser, err := FindUserByEmail(u.db, "me+testing2@julian.pro")
	assert.Equal(u.T(), err, nil)
	assert.NotNil(u.T(), foundUser.ID)

	byID, err := FindUserByID(u.db, foundUser.ID)
	assert.Equal(u.T(), err, nil)
	assert.NotNil(u.T(), byID.ID)

	users, err := FindUsers(u.db)
	assert.Equal(u.T(), err, nil)
	assert.NotEmpty(u.T(), users)
}

func TestUserTestSuite(t *testing.T) {
	config := &config.Configuration{
		DB: config.DBConfiguration{
			Driver: "sqlite",
			URL:    testDbName,
		},
	}
	db, err := storage.Dial(config)
	require.NoError(t, err)

	defer db.Close()
	db.AutoMigrate(&User{})

	ts := &UserTestSuite{
		db: db,
	}
	suite.Run(t, ts)

	// clean up db file from disk after test
	os.Remove(testDbName)
}

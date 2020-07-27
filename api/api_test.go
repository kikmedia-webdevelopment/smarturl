package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/models"
	"github.com/juliankoehn/mchurl/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	testDbName       = "database.db"
	linkCode         = "jpro"
	linkJSON         = `{"id":"jpro","url":"https://www.julian.pro","VisitCount":0}`
	linkUPDATEJSON   = `{"id":"jpro","url":"https://julian.pro","VisitCount":0}`
	testLoginJSON    = `{"email":"me+testing@julian.pro","password":"testpassword"}`
	testUserPassword = "testpassword"
	testUserEmail    = "me+testing@julian.pro"
)

type LinkTestSuite struct {
	suite.Suite
	api *API
}

func (l *LinkTestSuite) TestLinkA() {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(linkJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/links")

	if assert.NoError(l.T(), l.api.LinkCreate(c)) {
		assert.Equal(l.T(), 200, rec.Code)
	}
}

// Testing LinksList()
func (l *LinkTestSuite) TestLinkB() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/links")

	if assert.NoError(l.T(), l.api.LinksList(c)) {
		assert.Equal(l.T(), 200, rec.Code)
	}
}

// Testing loadLink()
func (l *LinkTestSuite) TestLinkC() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/links")
	c.SetParamNames("id")
	c.SetParamValues(linkCode)

	if assert.NoError(l.T(), l.api.loadLink(c)) {
		assert.Equal(l.T(), 301, rec.Code)
	}
}

// Testing LinkUpdate()
func (l *LinkTestSuite) TestLinkD() {
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(linkUPDATEJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/links")

	if assert.NoError(l.T(), l.api.LinkUpdate(c)) {
		assert.Equal(l.T(), 200, rec.Code)
	}
}

// Testing LinkDelete()
func (l *LinkTestSuite) TestLinkE() {
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(linkUPDATEJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/links")

	if assert.NoError(l.T(), l.api.LinkDelete(c)) {
		assert.Equal(l.T(), 204, rec.Code)
	}
}

// Creates a new Testing User and login
func (l *LinkTestSuite) TestUserA() {
	db := l.api.db
	user, password, err := models.NewUser(db, l.api.config, testUserEmail, testUserPassword)
	require.NoError(l.T(), err)
	assert.Equal(l.T(), password, testUserPassword)
	assert.Equal(l.T(), user.Email, testUserEmail)
	// create a new User

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testLoginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := l.api.echo.NewContext(req, rec)
	c.SetPath("/api/users/authenticate")

	if assert.NoError(l.T(), l.api.Login(c)) {
		assert.Equal(l.T(), 200, rec.Code)
	}

	var pair LoginResponse
	err = json.NewDecoder(rec.Body).Decode(&pair)
	require.NoError(l.T(), err)
	assert.NotEmpty(l.T(), pair.RefreshToken)

	requestString := fmt.Sprintf("{refresh_token: %s}", pair.RefreshToken)
	tokenRequest := refreshTokenRequest{
		RefreshToken: pair.RefreshToken,
	}
	payload, err := json.Marshal(tokenRequest)
	require.NoError(l.T(), err)
	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = l.api.echo.NewContext(req, rec)
	c.SetPath("/api/users/refresh")
	if assert.NoError(l.T(), l.api.RefreshToken(c)) {
		assert.Equal(l.T(), 200, rec.Code)
	}

	// test refresh
	l.T().Log(requestString)
}

func TestLinkSuite(t *testing.T) {
	config := &config.Configuration{
		DB: config.DBConfiguration{
			Driver: "sqlite",
			URL:    testDbName,
		},
	}
	db, err := storage.Dial(config)
	require.NoError(t, err)

	db.AutoMigrate(
		&models.User{},
		&models.AuditLogEntry{},
		&models.Link{},
	)

	api := New(db, config)

	ts := &LinkTestSuite{
		api: api,
	}
	suite.Run(t, ts)

	os.Remove(testDbName)
}

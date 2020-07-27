package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetQrCode(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:target")
	c.SetParamNames("target")
	c.SetParamValues("www.example.com")
	h := &API{}

	// Assertions
	if assert.NoError(t, h.GetQRCode(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

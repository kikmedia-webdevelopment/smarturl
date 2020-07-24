package api

import (
	"net/http"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/stores"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type API struct {
	store *stores.Store
	echo  *echo.Echo
}

// New starts a new Web-Service
func New(store *stores.Store, config *config.Configuration) {
	e := echo.New()
	e.HideBanner = true

	if config.Web.Debug {
		e.Debug = true
	}

	api := &API{
		store: store,
		echo:  e,
	}

	e.GET("/:id", api.getEntry)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func (a *API) getEntry(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusNotFound, "id not found")
	}

	entry, err := a.store.GetEntryAndIncrease(id)
	if err != nil {
		logrus.Error(err)
		return c.String(http.StatusNotFound, "id not found")
	}

	if entry.URL != "" {
		return c.Redirect(http.StatusMovedPermanently, entry.URL)
	}

	return c.String(http.StatusNotFound, "id not found")
}

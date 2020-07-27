package api

import (
	"github.com/juliankoehn/mchurl/models"
	"github.com/labstack/echo/v4"
)

func (a *API) listStats(c echo.Context) error {
	stats, err := models.ListStats(a.db)
	if err != nil {
		return err
	}
	return c.JSON(200, stats)
}

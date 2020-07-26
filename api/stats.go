package api

import "github.com/labstack/echo/v4"

func (a *API) listStats(c echo.Context) error {
	entriesCount, totalVisits, err := a.store.ListStats()
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"entries": entriesCount,
		"visits":  totalVisits,
	})
}

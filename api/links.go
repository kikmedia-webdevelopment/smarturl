package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/models"
	"github.com/labstack/echo/v4"
)

func (a *API) loadLink(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusNotFound, "id not found")
	}

	var link *models.Link

	err := a.db.Transaction(func(tx *gorm.DB) error {
		var terr error
		link, terr = models.GetLinkByID(tx, id)
		if terr != nil {
			return c.String(http.StatusNotFound, "id not found")
		}

		if terr := models.IncreaseVisitCounter(tx, link); terr != nil {
			return c.String(http.StatusNotFound, "id not found")
		}
		return nil
	})

	if err != nil {
		return c.String(http.StatusNotFound, "id not found")
	}

	if link.URL != "" {
		return c.Redirect(http.StatusMovedPermanently, link.URL)
	}

	return c.String(http.StatusNotFound, "id not found")
}

// LinkDelete deletes a single entry
func (a *API) LinkDelete(c echo.Context) error {
	l := new(models.Link)
	if err := c.Bind(l); err != nil {
		return err
	}
	if l.ID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing ID in Request")
	}

	if err := models.DeleteEntry(a.db, l.ID); err != nil {
		return err
	}
	return c.NoContent(204)
}

// LinkCreate creates a new Link
func (a *API) LinkCreate(c echo.Context) error {
	l := new(models.Link)
	if err := c.Bind(l); err != nil {
		return err
	}

	link, err := models.CreateEntry(a.db, a.config.DB.IDLength, l.URL, l.ID)
	if err != nil {
		return err
	}
	return c.JSON(200, link)
}

// LinksList lists all links
func (a *API) LinksList(c echo.Context) error {
	links, err := models.LinksList(a.db)
	if err != nil {
		// handle error
	}

	return c.JSON(200, links)
}

// LinkUpdate updates link url
func (a *API) LinkUpdate(c echo.Context) error {
	l := new(models.Link)
	if err := c.Bind(l); err != nil {
		return err
	}
	link, err := models.LinkUpdate(a.db, l)
	if err != nil {
		return err
	}
	return c.JSON(200, link)
}

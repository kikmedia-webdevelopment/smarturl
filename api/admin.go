package api

import (
	"html/template"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
)

func (a *API) serveAdmin(c echo.Context) error {
	buildPath := path.Clean("ui/build")
	tmpl, err := template.ParseFiles(path.Join("templates", "index.html"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing Template")
	}

	data, err := NewViewData(buildPath)
	data.Config.APIURL = a.config.Web.BaseURL
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing ViewData")
	}

	res := c.Response().Writer
	if err := tmpl.Execute(res, data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error serving Template")
	}
	return nil
}

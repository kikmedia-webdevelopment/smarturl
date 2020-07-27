package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/juliankoehn/mchurl/qrcode"
	"github.com/labstack/echo/v4"
)

// GetQRCode returns an QR Code for given URL
func (a *API) GetQRCode(c echo.Context) error {
	var dimension int
	var level qrcode.RecoveryLevel
	var target string
	var err error

	target = c.Param("target")
	if target == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Target")
	}

	target, err = url.QueryUnescape(target)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not decode URI")
	}

	dimensionString := c.Param("dimension")
	levelString := c.Param("level")
	if dimensionString != "" {
		dimension, err = strconv.Atoi(dimensionString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid dimension param, should be int")
		}
	} else {
		dimension = 256
	}

	if levelString != "" {
		lvl, err := strconv.Atoi(levelString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid level param, should be int")
		}

		// level low will always be transformed to level 1
		if (lvl >= 0) && (lvl <= 3) {
			switch lvl {
			case 0, 1:
				level = 1
			case 2:
				level = 2
			case 3:
				level = 3
			}
		} else {
			level = 1
		}

	} else {
		level = 1
	}

	png, err := qrcode.Encode(target, level, dimension)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, "image/png", png)
}

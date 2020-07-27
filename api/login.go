package api

import (
	"net/http"

	"github.com/juliankoehn/mchurl/models"
	"github.com/labstack/echo/v4"
)

type LoginResponse struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login a user
func (a *API) Login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Email")
	}

	usr, err := models.FindUserByEmail(a.db, u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if !usr.Authenticate(u.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Your given Password is Invalid")
	}

	pair, err := a.generateTokenPair(usr)
	if err != nil {
		return err
	}
	return c.JSON(200, pair)
}

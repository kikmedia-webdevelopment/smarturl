package api

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/juliankoehn/mchurl/models"
	"github.com/labstack/echo/v4"
)

type changePasswordParams struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

func (a *API) changePassword(c echo.Context) error {
	Token := c.Get("user").(*jwt.Token)
	claims := Token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	u := new(changePasswordParams)
	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := models.FindUserByEmail(a.db, email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "could not find user")
	}
	if user.Email != email {
		return echo.NewHTTPError(http.StatusNotFound, "could not find user")
	}

	if user.Authenticate(u.CurrentPassword) {
		if err := user.UpdatePassword(a.db, u.NewPassword); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error updating password")
		}
		return nil
	}

	return echo.NewHTTPError(http.StatusBadRequest, "current password is not valid")
}

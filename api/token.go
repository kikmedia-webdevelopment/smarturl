package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/juliankoehn/mchurl/models"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

// utc life

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type tokenPair struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *API) generateTokenPair(usr *models.User) (*tokenPair, error) {
	secret := a.config.JWT.Secret

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = usr.ID
	claims["email"] = usr.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Generate encoded token and send it as response
	tkn, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	ksuid := ksuid.New()
	refreshTokenString := ksuid.String()

	pair := &tokenPair{
		Token:        tkn,
		RefreshToken: refreshTokenString,
	}

	err = a.store.UserUpdateToken(usr.ID, refreshTokenString)
	if err != nil {
		return nil, err
	}

	return pair, nil
}

// RefreshToken handels token refreshs!
func (a *API) RefreshToken(c echo.Context) error {
	rtr := new(refreshTokenRequest)
	if err := c.Bind(rtr); err != nil {
		return err
	}

	if rtr.RefreshToken == "" {
		return echo.ErrUnauthorized
	}

	// user with token from database
	user, err := a.store.FindUserByToken(rtr.RefreshToken)
	if err != nil {
		return echo.ErrUnauthorized
	}

	if user != nil {
		now := time.Now()

		if user.TokenExpires.Before(now) {
			return echo.ErrUnauthorized
		}

		pair, err := a.generateTokenPair(user)
		if err != nil {
			return err
		}
		return c.JSON(200, pair)

	} else {
		return echo.ErrUnauthorized
	}
}

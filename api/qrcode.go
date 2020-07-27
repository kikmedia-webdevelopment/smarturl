package api

import (
	"net/http"

	"github.com/juliankoehn/mchurl/qrcode"
	"github.com/labstack/echo/v4"
)

// GetQRCodeParams
// level:
// low: 0, // Level L: 7% error recovery.
// medium: 1 // Level M: 15% error recovery. Good default choice.
// high: 2 // Level Q: 25% error recovery.
// highest: 3 // Level H: 30% error recovery.
type GetQRCodeParams struct {
	Size  int                  `json:"size"`
	Level qrcode.RecoveryLevel `json:"level"`
}

// GetQRCode returns an QR Code for given URL
func (a *API) GetQRCode(c echo.Context) error {
	params := new(GetQRCodeParams)
	if err := c.Bind(params); err != nil {
		// ignore error, as optional
		return nil
	}

	// level low will always be transformed to level 1
	if (params.Level >= 0) && (params.Level <= 3) {
		// nothing
	} else {
		params.Level = 1
	}

	if params.Level == 0 {
		params.Level = 1
	}

	if params.Size == 0 {
		// defaults to 256
		params.Size = 256
	}

	target := c.Param("target")
	if target == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Target")
	}

	png, err := qrcode.Encode(target, params.Level, params.Size)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, "image/png", png)
}

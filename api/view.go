package api

import (
	"fmt"

	"github.com/juliankoehn/mchurl/webpack"
)

// ViewData contains data for the view
type ViewData struct {
	Webpack *webpack.Webpack
	Config  struct {
		APIURL string `json:"apiUrl"`
	}
}

// NewViewData creates new data for the view
func NewViewData(buildPath string) (ViewData, error) {
	wp, err := webpack.New(buildPath)
	if err != nil {
		return ViewData{}, fmt.Errorf("failed to read webpack configuration: %w", err)
	}

	return ViewData{
		Webpack: wp,
	}, nil
}

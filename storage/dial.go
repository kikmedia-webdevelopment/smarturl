package storage

import (
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // default import for sqlite support
	"github.com/juliankoehn/mchurl/config"
	"github.com/pkg/errors"
)

// Connection is the interface a storage provider must implement.
type Connection struct {
	*gorm.DB
}

// Dial will connect to that storage engine
func Dial(config *config.Configuration) (*Connection, error) {
	if config.DB.Driver == "" && config.DB.URL != "" {
		u, err := url.Parse(config.DB.URL)
		if err != nil {
			return nil, errors.Wrap(err, "parsing db connection url")
		}
		config.DB.Driver = u.Scheme
	}
	if config.DB.Driver == "sqlite" {
		config.DB.Driver = "sqlite3"
	}

	db, err := gorm.Open(config.DB.Driver, config.DB.URL)
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}

	return &Connection{db}, nil
}

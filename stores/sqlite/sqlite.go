package sqlite

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/juliankoehn/mchurl/models"
	"github.com/juliankoehn/mchurl/stores/shared"
	"github.com/pkg/errors"
)

// BoltStore implements the stores.Storage interface
type SqliteStore struct {
	db *gorm.DB
}

func New(path string) (*SqliteStore, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "could not open sqlite3 database")
	}

	db.AutoMigrate(
		&shared.Entry{},
		&models.User{},
	)
	return &SqliteStore{
		db: db,
	}, nil
}

// Close closes the bolt database
func (s *SqliteStore) Close() error {
	return s.db.Close()
}

// CreateEntry creates an entry by a given ID and returns an error
func (s *SqliteStore) CreateEntry(entry shared.Entry, id string) (*shared.Entry, error) {
	currentTime := time.Now()
	entry.CreatedOn = &currentTime
	entry.ID = id

	if err := s.db.Create(&entry).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return nil, errors.New("entry already exists")
		}
		return nil, err
	}
	return &entry, nil
}

// DeleteEntry deleted an entry by a given ID and returns an error
func (s *SqliteStore) DeleteEntry(id string) error {
	if err := s.db.Where("id = ?", id).Delete(shared.Entry{}).Error; err != nil {
		return err
	}
	return nil
}

// GetEntryByID returns a entry and an error by the shorted ID
func (s *SqliteStore) GetEntryByID(id string) (*shared.Entry, error) {
	var entry shared.Entry
	if err := s.db.Where("id = ?", id).First(&entry).Error; err != nil {
		return nil, err
	}
	return &entry, nil
}

// IncreaseVisitCounter increases the visit counter and sets the current
// time as the last visit ones
func (s *SqliteStore) IncreaseVisitCounter(id string) error {
	if err := s.db.Model(shared.Entry{}).Where("id = ?", id).UpdateColumn("visit_count", gorm.Expr("visit_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// GetVisitors returns the visitors and an error of an entry
func (s *SqliteStore) GetVisitors(id string) ([]shared.Visitor, error) {
	return nil, nil
}

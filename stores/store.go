package stores

import (
	"strings"
	"time"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/models"
	"github.com/juliankoehn/mchurl/stores/shared"
	"github.com/juliankoehn/mchurl/stores/sqlite"
	"github.com/juliankoehn/mchurl/utils"
	"github.com/pkg/errors"
)

// Store holds internal funcs and vars about the store
type Store struct {
	storage  shared.Storage
	idLength int
}

// ErrNoValidURL is returned when the URL is not valid
var ErrNoValidURL = errors.New("the given URL is no valid URL")

// ErrGeneratingIDFailed is returned when the 10 tries to generate an id failed
var ErrGeneratingIDFailed = errors.New("could not generate unique id, all ten tries failed")

// ErrEntryIsExpired is returned when the entry is expired
var ErrEntryIsExpired = errors.New("entry is expired")

// New starts a new storage
func New(config *config.DBConfiguration) (*Store, error) {
	if config.IDLength == 0 {
		config.IDLength = 4
	}
	var err error
	var s shared.Storage

	if config.URL == "" {
		return nil, errors.New("missing database url")
	}

	switch driver := config.Driver; driver {
	case "sqlite":
		s, err = sqlite.New(config.URL)
	default:
		return nil, errors.New(driver + " is not a recognized database driver")
	}

	if err != nil {
		return nil, errors.Wrap(err, "could not initialize the database")
	}

	return &Store{
		storage:  s,
		idLength: config.IDLength,
	}, nil
}

func (s *Store) CreateEntry(entry shared.Entry, givenID string) (*shared.Entry, error) {
	entry.URL = strings.Replace(entry.URL, " ", "%20", -1)
	var id string
	var err error

	if givenID != "" {
		id = givenID
	} else {
		id, err = utils.GenerateRandomString(s.idLength)
		if err != nil {
			return nil, err
		}
	}

	link, err := s.storage.CreateEntry(entry, id)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// GetEntryByID returns an entry by ID
func (s *Store) GetEntryByID(id string) (*shared.Entry, error) {
	if id == "" {
		return nil, shared.ErrNoEntryFound
	}
	return s.storage.GetEntryByID(id)
}

// GetEntryAndIncrease the visitor count
//
func (s *Store) GetEntryAndIncrease(id string) (*shared.Entry, error) {
	entry, err := s.GetEntryByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch entry "+id)
	}
	if entry.Expiration != nil && !entry.Expiration.IsZero() && time.Now().After(*entry.Expiration) {
		return nil, ErrEntryIsExpired
	}
	if err := s.storage.IncreaseVisitCounter(id); err != nil {
		return nil, errors.Wrap(err, "could not increase visitor counter")
	}

	entry.VisitCount++
	return entry, nil
}

// DeleteEntry deletes an Entry fully from the DB
func (s *Store) DeleteEntry(id string) error {
	return s.storage.DeleteEntry(id)
}

// CreateUser creates a new User
func (s *Store) CreateUser(user models.User) (*models.User, error) {
	return s.storage.CreateUser(user)
}

// FindUserByEmail returns a user if given by email
func (s *Store) FindUserByEmail(email string) (*models.User, error) {
	return s.storage.FindUserByEmail(email)
}

// LinksList lists all available links
func (s *Store) LinksList() ([]*shared.Entry, error) {
	return s.storage.LinksList()
}

// LinkUpdate updates a link
func (s *Store) LinkUpdate(link *shared.Entry) (*shared.Entry, error) {
	return s.storage.LinkUpdate(link)
}

func (s *Store) UserUpdateToken(id uint, token string) error {
	return s.storage.UserUpdateToken(id, token)
}

func (s *Store) FindUserByToken(token string) (*models.User, error) {
	return s.storage.FindUserByToken(token)
}

package shared

import (
	"errors"
	"time"

	"github.com/juliankoehn/mchurl/models"
)

// Storage is an interface which will be implmented by each storage
// e.g. bolt, sqlite
type Storage interface {
	GetEntryByID(string) (*Entry, error)
	GetVisitors(string) ([]Visitor, error)
	DeleteEntry(string) error
	IncreaseVisitCounter(*Entry) error
	CreateEntry(Entry, string) (*Entry, error)
	CreateUser(models.User) (*models.User, error)
	FindUserByEmail(string) (*models.User, error)
	FindUserByToken(token string) (*models.User, error)
	LinksList() ([]*Entry, error)
	LinkUpdate(*Entry) (*Entry, error)
	UserUpdateToken(id uint, token string) error
	Close() error

	// stats
	ListStats() (*int, *int, error)
}

// Entry is the data set which is stored in the DB as JSON
type Entry struct {
	// RemoteAddr is the clients address
	ID                    string     `json:"id"`
	URL                   string     `json:"url"`
	RemoteAddr            string     `json:",omitempty"`
	DeletionURL           string     `json:",omitempty"`
	Password              []byte     `json:",omitempty"`
	LastVisit, Expiration *time.Time `json:",omitempty"`
	CreatedOn             *time.Time `json:"-"`
	VisitCount            int
}

// Visitor is the entry which is stored in the visitors bucket
type Visitor struct {
	IP, Referer, UserAgent                                 string
	Timestamp                                              time.Time
	UTMSource, UTMMedium, UTMCampaign, UTMContent, UTMTerm string `json:",omitempty"`
}

// ErrNoEntryFound is returned when no entry to a id is found
var ErrNoEntryFound = errors.New("no entry found with this ID")

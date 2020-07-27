package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/utils"
)

// Link is the data for our "relinker-url-shortener"
type Link struct {
	ID                    string     `json:"id"`
	URL                   string     `json:"url"`
	RemoteAddr            string     `json:",omitempty"`
	DeletionURL           string     `json:",omitempty"`
	Password              []byte     `json:",omitempty"`
	LastVisit, Expiration *time.Time `json:",omitempty"`
	CreatedOn             *time.Time `json:"-"`
	VisitCount            int
}

// CreateEntry  creates an entry by a given ID and returns an error
func CreateEntry(tx *gorm.DB, length int, url, id string) (*Link, error) {
	var err error
	currentTime := time.Now()

	link := Link{
		CreatedOn: &currentTime,
	}
	link.URL = strings.Replace(url, " ", "%20", -1)
	if id == "" {
		id, err = utils.GenerateRandomString(length)
		if err != nil {
			return nil, err
		}
	}
	link.ID = id

	if err := tx.Create(&link).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return nil, errors.New("entry already exists")
		}
		return nil, err
	}
	return &link, nil
}

// DeleteEntry deleted an entry by a given ID and returns an error
func DeleteEntry(tx *gorm.DB, id string) error {
	if err := tx.Where("id = ?", id).Delete(Link{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

// GetLinkByID returns a entry and an error by the shorted ID
func GetLinkByID(tx *gorm.DB, id string) (*Link, error) {
	var link Link
	if err := tx.Where("id = ?", id).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// IncreaseVisitCounter increases the visit counter and sets the current
func IncreaseVisitCounter(tx *gorm.DB, link *Link) error {
	currentTime := time.Now()

	if err := tx.Model(Link{}).Where("id = ?", link.ID).Updates(map[string]interface{}{
		"visit_count": gorm.Expr("visit_count + ?", 1),
		"last_visit":  currentTime,
	}).Error; err != nil {
		return err
	}
	return nil
}

// LinksList returns a List of all available Links formerly "entries"
func LinksList(tx *gorm.DB) ([]*Link, error) {
	var links []*Link

	if err := tx.Find(&links).Error; err != nil {
		return nil, err
	}

	return links, nil
}

// LinkUpdate updates URL of Link
func LinkUpdate(tx *gorm.DB, link *Link) (*Link, error) {
	if err := tx.Model(&Link{}).Where("id = ?", link.ID).UpdateColumn("url", link.URL).Error; err != nil {
		return nil, err
	}
	return link, nil
}

package models

import (
	"github.com/jinzhu/gorm"
)

// Stats returns entries and visits for our dashboard view
type Stats struct {
	Entries int `json:"entries"`
	Visits  int `json:"visits"`
}

// ListStats lists active links and clicks
func ListStats(tx *gorm.DB) (*Stats, error) {
	links, err := LinksList(tx)
	if err != nil {
		return nil, err
	}
	var linksCount int
	var totalvisits int

	linksCount = len(links)
	for _, link := range links {
		totalvisits += link.VisitCount
	}

	return &Stats{
		Entries: linksCount,
		Visits:  totalvisits,
	}, nil
}

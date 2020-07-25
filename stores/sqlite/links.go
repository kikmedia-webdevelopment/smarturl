package sqlite

import "github.com/juliankoehn/mchurl/stores/shared"

// LinksList returns a List of all available Links formerly "entries"
func (s *SqliteStore) LinksList() ([]*shared.Entry, error) {
	var links []*shared.Entry

	if err := s.db.Find(&links).Error; err != nil {
		return nil, err
	}

	return links, nil
}

func (s *SqliteStore) LinkUpdate(link *shared.Entry) (*shared.Entry, error) {
	if err := s.db.Model(&shared.Entry{}).Where("id = ?", link.ID).UpdateColumn("url", link.URL).Error; err != nil {
		return nil, err
	}
	return link, nil
}

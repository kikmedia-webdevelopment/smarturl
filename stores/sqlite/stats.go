package sqlite

func (s *SqliteStore) ListStats() (*int, *int, error) {
	links, err := s.LinksList()
	if err != nil {
		return nil, nil, err
	}
	var linksCount int
	var totalvisits int

	linksCount = len(links)
	for _, link := range links {
		totalvisits += link.VisitCount
	}

	return &linksCount, &totalvisits, nil
}

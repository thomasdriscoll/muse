package services

type Scrapper interface {
	Scrape(source string, sourceType string) ([]byte, error)
}

type ScrapperImpl struct {
}

func wikiScrape(url string) []byte {
	return nil
}

func (s *ScrapperImpl) Scrape(source string, sourceType string) ([]byte, error) {
	return nil, nil
}

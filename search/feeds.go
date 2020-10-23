package search

import (
	"encoding/json"
	"log"
	"os"
)

type Search struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

var search = "datas/search.json"

// Get all feeds type to search, feeds in datas/search.json
func GetFeeds() ([]*Search, error) {
	searchsFile, err := os.Open(search)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer searchsFile.Close()

	var searchs []*Search

	err = json.NewDecoder(searchsFile).Decode(&searchs)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return searchs, nil
}

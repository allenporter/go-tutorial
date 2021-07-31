package search

import (
  "os"
  "encoding/json"
)

const feedsFile = "data/feeds.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(feedsFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}

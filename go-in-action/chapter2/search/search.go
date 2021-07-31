package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	matcher, exists := matchers["rss"]
	if !exists {
		matcher = matchers["default"]
	}

	for _, feed := range feeds {
		go func(feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

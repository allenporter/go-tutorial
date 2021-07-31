package main

import (
	"log"
	"os"

	_ "example.com/go-in-action/chapter2/matchers"
	"example.com/go-in-action/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("fake")
}

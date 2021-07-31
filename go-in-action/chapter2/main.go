package main

import (
    "log"
    "os"

    "example.com/go-in-action/chapter2/search"
  _ "example.com/go-in-action/chapter2/matchers"
)

func init() {
  log.SetOutput(os.Stdout)
}

func main() {
  search.Run("fake")
}

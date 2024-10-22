package main

import (
	"encoding/json"
	"os"

	"github.com/jhrick/quotes-from-thinkers/internal"
)

var quoteService internal.QuoteService

func main() {
  quoteChannel := make(chan internal.QuoteSchema)
  defer close(quoteChannel)

  quoteService = internal.QuoteService{
    QuoteChannel: quoteChannel,
  }

  go quoteService.GetData()

  stdout := json.NewEncoder(os.Stdout)
  stdout.SetIndent("", "  ")

  for quote := range quoteService.QuoteChannel {
    stdout.Encode(internal.QuoteSchema{Author: quote.Author, Text: quote.Text})
  }
}

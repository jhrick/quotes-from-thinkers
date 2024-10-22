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

  subdirectory := "/frases_pensadores/1"

  go quoteService.GetData(subdirectory)

  stdout := json.NewEncoder(os.Stdout)
  stdout.SetIndent("", "  ")

  for quote := range quoteService.QuoteChannel {
    stdout.Encode(internal.QuoteSchema{Author: quote.Author, Text: quote.Text})
  }
}

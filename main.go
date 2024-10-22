package main

import (
	"encoding/json"
	"os"

	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

var quoteService services.QuoteService

func main() {
  quoteChannel := make(chan services.QuoteSchema)
  defer close(quoteChannel)

  scrapperService := services.ScrapperService(quoteChannel)

  quoteService = services.QuoteService{
    QuoteChannel: quoteChannel,
  }

  subdirectory := "/frases_pensadores/1"

  go scrapperService.GetData(subdirectory)

  stdout := json.NewEncoder(os.Stdout)
  stdout.SetIndent("", "  ")

  for quote := range quoteService.QuoteChannel {
    stdout.Encode(services.QuoteSchema{Author: quote.Author, Text: quote.Text})
  }
}

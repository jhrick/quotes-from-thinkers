package main

import (
	"encoding/json"
	"os"

	"github.com/jhrick/quotes-from-thinkers/internal/services"
)


func main() {
  quoteChannel := make(chan services.QuoteSchema)

  scrapperService := services.ScrapperService(quoteChannel)

  quoteService := services.QuoteService(scrapperService.QuoteChannel)

  subdirectory := "/frases_pensadores/1"

  go scrapperService.GetData(subdirectory, 2)

  quoteService.GetQuotes()
  
  stdout := json.NewEncoder(os.Stdout)
  stdout.SetIndent("", "  ")

  for quote := range quoteService.QuoteChannel {
    stdout.Encode(services.QuoteSchema{Author: quote.Author, Text: quote.Text})
  }
}

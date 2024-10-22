package main

import "github.com/jhrick/quotes-from-thinkers/internal/services"

func main() {
  quoteChannel := make(chan services.QuoteSchema)

  scrapperService := services.ScrapperService(quoteChannel)

  quoteService := services.QuoteService(scrapperService.QuoteChannel)

  subdirectory := "/frases_pensadores/1"

  go scrapperService.GetData(subdirectory, 2)

  quoteService.GetQuotes()
}

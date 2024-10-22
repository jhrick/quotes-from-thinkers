package services

import (
	"encoding/json"
	"log"
	"os"
)

type QuoteSchema struct {
  Author string
  Text   string
}

type implQuote struct {
  QuoteChannel chan QuoteSchema
}

func QuoteService(quoteChannel chan QuoteSchema) implQuote {
  impl := &implQuote{
    QuoteChannel: quoteChannel,
  }

  return *impl
}

var quotes map[string][]string

func (q *implQuote) GetQuotes() {
  fName := "quotes.json"
  file, err := os.Create(fName)
  if err != nil {
    log.Fatal(err.Error())
    return
  }
  defer file.Close()

  quotes = make(map[string][]string)

  for quote := range q.QuoteChannel {
    quotes[quote.Author] = append(quotes[quote.Author], quote.Text)

    _, ok := <-q.QuoteChannel
    if !ok {
      break
    }
  }

  enc := json.NewEncoder(file)
  enc.SetIndent("", "  ")

  enc.Encode(quotes)
}
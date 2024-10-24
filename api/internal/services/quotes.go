package services

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/jhrick/quotes-from-thinkers/internal/repository"
)


type QuotesSchema struct {
  ID     string
  Author string
  Text   string
}

type QuotesModel struct {
  ID   string
  Text string
}

type implQuotes struct {
  QuotesChannel chan QuotesSchema
}

func QuotesService(quotesChannel chan QuotesSchema) implQuotes {
  impl := &implQuotes{
    QuotesChannel: quotesChannel,
  }

  return *impl
}

var quotes map[string][]QuotesModel

func (q *implQuotes) InsertQuotes(wg *sync.WaitGroup) {
  authors := make(map[string]string)

  for quote := range q.QuotesChannel {
    _, ok := <-q.QuotesChannel
    if !ok {
      break
    }

    wg.Add(1)

    go func() {
      defer wg.Done()

      authorId, ok := authors[quote.Author]
      if !ok {
        id, err := repository.Author.Create(quote.Author)
        if err != nil {
          panic(err)
        }

        authors[quote.Author] = id
        authorId = id
      }

      err := repository.Quotes.Create(quote.ID, authorId, quote.Text)
      if err != nil {
        panic(err)
      }
    }()
  }
}

func (q *implQuotes) GetQuotes() {
  fName := "quotes.json"
  file, err := os.Create(fName)
  if err != nil {
    log.Fatal(err.Error())
    return
  }
  defer file.Close()

  quotes = make(map[string][]QuotesModel)

  for quote := range q.QuotesChannel {
    line := QuotesModel{ ID: quote.ID, Text: quote.Text }
    quotes[quote.Author] = append(quotes[quote.Author], line)

    _, ok := <-q.QuotesChannel
    if !ok {
      break
    }
  }

  enc := json.NewEncoder(file)
  enc.SetIndent("", "  ")

  enc.Encode(quotes)
}

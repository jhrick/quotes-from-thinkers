package services

import (
	"log"
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

type ImplQuotes struct {}

func QuotesService() ImplQuotes {
  impl := &ImplQuotes{}

  return *impl
}

var quotes map[string][]QuotesModel

func (q *ImplQuotes) InsertQuotes(data []QuotesSchema, wg *sync.WaitGroup) {
  authors := make(map[string]string)

  for i := 0; i < len(data); i++ {
    wg.Add(1)

    quote := data[i]

    go func() {
      defer wg.Done()

      authorId, ok := authors[quote.Author]
      if !ok {
        id, err := repository.Author.Create(quote.Author)
        if err != nil {
          log.Println("quotes:", err)
          return
        }

        authors[quote.Author] = id
        authorId = id
      }

      err := repository.Quotes.Create(quote.ID, authorId, quote.Text)
      if err != nil {
        log.Println("quotes:", err)
        return
      }
    }()
  }
}

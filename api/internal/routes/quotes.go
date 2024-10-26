package routes

import (
	"net/http"
	"sync"

	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

func (a *apiHandler) handleImportQuotes(w http.ResponseWriter, _ *http.Request) {
  var wg sync.WaitGroup

  a.s.Quotes.InsertQuotes(&wg)

  wg.Wait()

  w.Write([]byte("finish"))
}

func (a *apiHandler) handleGetQuote(w http.ResponseWriter, _ *http.Request) {
  a.mu.Lock()
  defer a.mu.Unlock()

  quote := <-a.qChan

  type response struct {
    Quote services.QuotesSchema `json:"quote"`
  }

  sendJSON(w, response{Quote: quote})
}

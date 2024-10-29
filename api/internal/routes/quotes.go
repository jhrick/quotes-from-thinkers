package routes

import (
	"net/http"
	"sync"

	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

var quotesService services.ImplQuotes = services.QuotesService()

func (a *apiHandler) handleImportQuotes(w http.ResponseWriter, _ *http.Request) {
  var wg sync.WaitGroup

  wg.Wait()
}

func (a *apiHandler) handleGetQuote(w http.ResponseWriter, _ *http.Request) {
  a.mu.Lock()
  defer a.mu.Unlock()

  type response struct {
    Quote services.QuotesSchema `json:"quote"`
  }
}

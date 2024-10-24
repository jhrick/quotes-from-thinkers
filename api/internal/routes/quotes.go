package routes

import (
	"net/http"
	"sync"
)

func (a *apiHandler) handleImportQuotes(w http.ResponseWriter, _ *http.Request) {
  var wg sync.WaitGroup

  a.s.Quotes.InsertQuotes(&wg)

  wg.Wait()

  w.Write([]byte("finish"))
}

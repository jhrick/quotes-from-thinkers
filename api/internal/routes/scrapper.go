package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *apiHandler) handlerScrapper(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Connection", "Keep-Alive")
  w.Header().Set("Transfer-Encoding", "chunked")
  w.Header().Set("X-Content-Type-Options", "nosniff")

  flusher, ok := w.(http.Flusher)
  if !ok {
    w.WriteHeader(501)

    type failed struct {
      Message string `json:"message"`
    }

    sendJSON(w, failed{Message:"internal-server-error"})
  }

  limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))
  
  subdirectory := "/frases_pensadores/1"

  go a.s.Scrapper.GetData(subdirectory, limit)

  type response struct {
    Quote string `json:"quote"`
  }

  for quote := range a.qChan {
    sendJSON(w, response{Quote: quote.Text})

    flusher.Flush()
  }
}

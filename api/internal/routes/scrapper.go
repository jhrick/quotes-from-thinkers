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

  limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))
  
  subdirectory := "/frases_pensadores/1"

  go a.s.Scrapper.GetData(subdirectory, limit)

  type response struct {
    Success bool   `json:"success"`
    Error   error  `json:"error"`
  }

  if len(a.errChan) != 0 {
    w.WriteHeader(501)

    err := <-a.errChan

    sendJSON(w, response{Success: false, Error: err})
  }

  sendJSON(w, response{Success: true, Error: nil})
}

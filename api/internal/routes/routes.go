package routes

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type apiHandler struct {
  r        *chi.Mux
  upgrader websocket.Upgrader
  mu       *sync.Mutex
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  h.r.ServeHTTP(w, r)
}

func NewHandler() http.Handler {
  a := apiHandler{
    upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
    mu: &sync.Mutex{},
  }

  r := chi.NewRouter()

  r.Route("/api", func(r chi.Router) {
    r.Route("/scrapper/", func(r chi.Router) {
      r.Get("/", a.handlerScrapper)
    })
    r.Route("/quotes", func(r chi.Router) {
      r.Post("/import", a.handleImportQuotes)
      r.Get("/", a.handleGetQuote)
    })
  })

  a.r = r

  return a
}

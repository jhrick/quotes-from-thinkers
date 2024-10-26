package routes

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

type apiHandler struct {
  r       *chi.Mux
  s       services.ImplServices
  mu      *sync.Mutex
  qChan   chan services.QuotesSchema
  errChan chan error
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  h.r.ServeHTTP(w, r)
}

func NewHandler() http.Handler {
  quotesChannel := make(chan services.QuotesSchema)
  errChannel := make(chan error)

  a := apiHandler{
    s: services.Services(quotesChannel, errChannel),
    mu: &sync.Mutex{},
    qChan: quotesChannel,
    errChan: errChannel,
  }

  r := chi.NewRouter()

  r.Route("/api", func(r chi.Router) {
    r.Route("/scrapper/", func(r chi.Router) {
      r.Get("/{limit}", a.handlerScrapper)
    })
    r.Route("/quotes", func(r chi.Router) {
      r.Post("/", a.handleImportQuotes)
      r.Get("/", a.handleGetQuote)
    })
  })

  a.r = r

  return a
}

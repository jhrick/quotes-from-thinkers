package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

type apiHandler struct {
  r     *chi.Mux
  s     services.ImplServices
  qChan chan services.QuotesSchema
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  h.r.ServeHTTP(w, r)
}

func NewHandler(quotesChannel chan services.QuotesSchema) http.Handler {
  a := apiHandler{
    s: services.Services(quotesChannel),
    qChan: quotesChannel,
  }

  r := chi.NewRouter()

  r.Route("/api", func(r chi.Router) {
    r.Route("/scrapper/", func(r chi.Router) {
      r.Get("/{limit}", a.handlerScrapper)
    })
    r.Route("/quotes", func(r chi.Router) {
      r.Post("/", a.handleImportQuotes)
    })
  })

  a.r = r

  return a
}

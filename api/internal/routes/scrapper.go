package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jhrick/quotes-from-thinkers/internal/services"
)

var scrapperService services.ImplScrapper = services.ScrapperService()

func (a *apiHandler) handlerScrapper(w http.ResponseWriter, r *http.Request) {
  limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
  
  c, err := a.upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Print("upgrade:", err)
    return
  }
  defer c.Close()

  quotesChannel := make(chan services.QuotesSchema)

  subdirectory := "/frases_pensadores/1"

  go scrapperService.GetData(quotesChannel, subdirectory, limit)

  type response struct {
    Success bool `json:"success"`
  }

  for quotes := range quotesChannel {
    if err := c.WriteJSON(quotes); err != nil {
      log.Println("write:", err)
      break
    }
  }
}

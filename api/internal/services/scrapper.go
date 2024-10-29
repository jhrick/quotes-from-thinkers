package services

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type ImplScrapper struct {}

func ScrapperService() ImplScrapper {
  impl := &ImplScrapper{}

  return *impl
}

func (s *ImplScrapper) GetData(quotesChannel chan QuotesSchema, subdirectory string, limit int) {
  domain := "https://www.pensador.com"

  c := colly.NewCollector()

  c.OnError(func(r *colly.Response, err error) {
    log.Println("scapper:", err)
  })

  c.OnHTML("div.thought-card", func (e *colly.HTMLElement) {
    id := e.ChildAttr("p", "id")
    author := e.ChildText("span.author-name")
    text := e.ChildText("p.frase")

    quote := QuotesSchema{
      ID: id,
      Author: author,
      Text: text,
    }
 
    quotesChannel <- quote
  })

  c.OnHTML("a.nav", func (e *colly.HTMLElement) {
    if !strings.Contains(e.Text, "Próxima") {
      return
    }

    pageNum := getPageNum(subdirectory)

    if pageNum >= strconv.Itoa(limit) {
      return
    }

    subdirectory = e.Attr("href")

    s.GetData(quotesChannel, subdirectory, limit)
  })

  link := fmt.Sprintf("%s%s", domain, subdirectory)

  fmt.Println(link)

  c.Visit(link)
}

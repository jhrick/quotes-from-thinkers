package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type implScrapper struct {
  quotesChannel chan QuotesSchema
  errChannel    chan error
}

func ScrapperService(quoteChannel chan QuotesSchema, errChannel chan error) implScrapper {
  impl := &implScrapper{
    quotesChannel: quoteChannel,
    errChannel: errChannel,
  }

  return *impl
}

func (s *implScrapper) GetData(subdirectory string, limit int) {
  domain := "https://www.pensador.com"

  c := colly.NewCollector()

  c.OnError(func(r *colly.Response, err error) {
    s.errChannel <- err
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
 
    s.quotesChannel <- quote
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

    s.GetData(subdirectory, limit)
  })

  link := fmt.Sprintf("%s%s", domain, subdirectory)

  fmt.Println(link)

  c.Visit(link)
}

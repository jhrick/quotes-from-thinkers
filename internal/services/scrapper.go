package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type implScrapper struct {
  QuoteChannel chan QuoteSchema
}

func ScrapperService(quoteChannel chan QuoteSchema) implScrapper {
  impl := &implScrapper{
    QuoteChannel: quoteChannel,
  }

  return *impl
}

func (s *implScrapper) GetData(subdirectory string, limit int) {
  domain := "https://www.pensador.com"

  c := colly.NewCollector()

  c.OnHTML("div.thought-card", func (e *colly.HTMLElement) {
    id := e.ChildAttr("p", "id")
    author := e.ChildText("span.author-name")
    text := e.ChildText("p.frase")

    quote := QuoteSchema{
      ID: id,
      Author: author,
      Text: text,
    }

    s.QuoteChannel <- quote
  })

  c.OnHTML("a.nav", func (e *colly.HTMLElement) {
    if !strings.Contains(e.Text, "Próxima") {
      return
    }

    pageNum := getPageNum(subdirectory)

    if pageNum >= strconv.Itoa(limit) {
      close(s.QuoteChannel)
      return
    }

    subdirectory = e.Attr("href")

    s.GetData(subdirectory, limit)
  })

  link := fmt.Sprintf("%s%s", domain, subdirectory)

  fmt.Println(link)

  c.Visit(link)
}

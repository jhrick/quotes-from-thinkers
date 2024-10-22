package services

import (
	"fmt"
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

func (s *implScrapper) GetData(subdirectory string) {
  domain := "https://www.pensador.com"

  c := colly.NewCollector()

  c.OnHTML("div.thought-card", func (e *colly.HTMLElement) {
    author := e.ChildText("span.author-name")
    text := e.ChildText("p.frase.fr")

    quote := QuoteSchema{
      Author: author,
      Text: text,
    }

    s.QuoteChannel <- quote
  })

  c.OnHTML("a.nav", func (e *colly.HTMLElement) {
    if !strings.Contains(e.Text, "Próxima") {
      return
    }

    subdirectory = e.Attr("href")

    s.GetData(subdirectory)
  })

  link := fmt.Sprintf("%s%s", domain, subdirectory)

  fmt.Println(link)

  c.Visit(link)
}

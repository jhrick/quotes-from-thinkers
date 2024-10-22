package internal

import "github.com/gocolly/colly/v2"

type QuoteSchema struct {
  Author string
  Text   string
}

type QuoteService struct {
  QuoteChannel chan QuoteSchema
}

func (q *QuoteService) GetData() {
  c := colly.NewCollector()

  link := "https://www.pensador.com/frases_pensadores/"

  c.OnHTML("div.thought-card", func (e *colly.HTMLElement) {
    author := e.ChildText("span.author-name")
    text := e.ChildText("p.frase.fr")

    quote := QuoteSchema{
      Author: author,
      Text: text,
    }

    q.QuoteChannel <- quote
  })

  c.Visit(link)
}

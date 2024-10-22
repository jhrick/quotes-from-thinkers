package services

type QuoteSchema struct {
  Author string
  Text   string
}

type implQuote struct {
  QuoteChannel chan QuoteSchema
}

func QuotesService(quoteChannel chan QuoteSchema) implQuote {
  impl := &implQuote{
    QuoteChannel: quoteChannel,
  }

  return *impl
}

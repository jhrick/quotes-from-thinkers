package services

type ImplServices struct {
  Scrapper implScrapper
  Quotes   implQuotes
}

func Services(quotesChannel chan QuotesSchema, errChannel chan error) ImplServices {
  implScrapper := ScrapperService(quotesChannel, errChannel)
  implQuotes := QuotesService(quotesChannel, errChannel)

  impl := ImplServices{
    Scrapper: implScrapper,
    Quotes: implQuotes,
  }

  return impl
}

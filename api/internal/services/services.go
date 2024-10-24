package services

type ImplServices struct {
  Scrapper implScrapper
  Quotes   implQuotes
}

func Services(quotesChannel chan QuotesSchema) ImplServices {
  implScrapper := ScrapperService(quotesChannel)
  implQuotes := QuotesService(quotesChannel)

  impl := ImplServices{
    Scrapper: implScrapper,
    Quotes: implQuotes,
  }

  return impl
}

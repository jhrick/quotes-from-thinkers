package repository

import (
	"github.com/jackc/pgx/v5/pgconn"
)

type implQuotes struct {}

func quotesRepo() implQuotes {
  impl := implQuotes{}

  return impl
}

func (q * implQuotes) Create(id string, authorId string, text string) *pgconn.PgError {
  query := `INSERT INTO "quotes" ("id", "author_id", "text") VALUES ($1, $2, $3)`

  err := pool.QueryRow(
    ctx,
    query,
    id,
    authorId,
    text,
  ).Scan()

  if err != nil {
    return parsePgError(err)
  }

  return nil
}

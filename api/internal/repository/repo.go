package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jhrick/quotes-from-thinkers/internal/store"
)

var pool *pgxpool.Pool = store.Pool

var ctx context.Context = store.Ctx 

var (
  Author = authorRepo()
  Quotes = quotesRepo()
)

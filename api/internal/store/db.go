package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
  Pool *pgxpool.Pool
  Ctx context.Context
)

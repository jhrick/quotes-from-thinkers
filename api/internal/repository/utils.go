package repository

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func parsePgError(err error) *pgconn.PgError {
  var pgErr *pgconn.PgError

  if errors.As(err, &pgErr) {
    return pgErr
  }

  return nil
}


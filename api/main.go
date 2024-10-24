package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jhrick/quotes-from-thinkers/internal/routes"
	"github.com/jhrick/quotes-from-thinkers/internal/services"
	"github.com/jhrick/quotes-from-thinkers/internal/store"
	"github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load("../.env"); err != nil {
    panic(err)
  }

  ctx := context.Background()

  pool, err := pgxpool.New(ctx, fmt.Sprintf(
    "user=%s password=%s host=%s port=%s dbname=%s",
    os.Getenv("DATABASE_USER"),
    os.Getenv("DATABASE_PASSWORD"),
    os.Getenv("DATABASE_HOST"),
    os.Getenv("DATABASE_PORT"),
    os.Getenv("DATABASE_NAME"),
  ))
  if err != nil {
    panic(err)
  }

  defer pool.Close()

  if err := pool.Ping(ctx); err != nil {
    panic(err)
  }

  store.Pool = pool
  store.Ctx = ctx
  
  quotesChannel := make(chan services.QuotesSchema)

  handler := routes.NewHandler(quotesChannel)

  go func() {
    if err := http.ListenAndServe(":8080", handler); err != nil {
      if !errors.Is(err, http.ErrServerClosed) {
        panic(err)
      }
    }
  }()

  log.Println("Server Running!")

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)
  <-quit
}

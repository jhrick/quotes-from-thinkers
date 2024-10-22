package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load("../.env"); err != nil {
    panic(err)
  }

  output, err := exec.Command(
    "tern",
    "migrate",
    "--migrations",
    "./internal/store/migrations",
    "--config",
    "./internal/store/migrations/tern.conf",
  ).Output()
  
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s:%s\n", "tern", string(output))
}

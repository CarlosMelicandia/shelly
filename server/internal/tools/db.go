package tools

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/weareinit/Opal/internal/config"
)

type LoadDBFunc func(db *sql.DB) error

func LoadDB(operation LoadDBFunc) error {
  envConfig := config.LoadEnv()

  url := fmt.Sprintf("libsql://%s?authToken=%s", envConfig.TursoDatabaseName, envConfig.TursoAuthToken)

  db, err := sql.Open("libsql", url)
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
    os.Exit(1)
  }

  if err := db.Ping(); err != nil {
      return fmt.Errorf("failed to connect to the database: %v", err)
  }

  defer db.Close()

  if err := operation(db); err != nil {
      fmt.Fprintf(os.Stderr, "operation failed: %s\n", err)
  }

  return nil
}

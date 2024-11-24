package tools

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/weareinit/Opal/internal/config"
)

type LoadDBFunc[T any] func(db *sql.DB) (T, error)

func LoadDB[T any](operation LoadDBFunc[T]) (T, error) {
	envConfig := config.LoadEnv()

	url := fmt.Sprintf("libsql://%s?authToken=%s", envConfig.TursoDatabaseName, envConfig.TursoAuthToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		var empty T
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
		return empty, err
	}

	if err := db.Ping(); err != nil {
		var empty T
		return empty, fmt.Errorf("failed to connect to the database: %v", err)
	}

	defer db.Close()

	result, err := operation(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "operation failed: %s\n", err)
		return result, err
	}

	return result, nil
}

package integration

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPostgresIntegration(t *testing.T) {
	ctx := context.Background()

	// 1. Setup Container (Conceptual/Snippet)
	// Catatan: Diperlukan 'go get' untuk dependensi testcontainers
	dbName := "testdb"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)
	if err != nil {
		t.Skip("Docker missing or Testcontainers not reachable. Skipping integration test.")
		return
	}

	// 2. Cleanup Otomatis
	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// 3. Get Connection String
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %s", err)
	}

	// 4. Test DB Connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("failed to open db: %s", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("failed to ping db: %s", err)
	}

	fmt.Println("Successfully connected to ephemeral Postgres container!")
}

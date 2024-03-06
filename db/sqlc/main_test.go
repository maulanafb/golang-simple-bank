package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatal("cannot begin transaction:", err)
	}

	// Assign the testQueries with the transaction
	testQueries = New(tx)

	// Run tests
	exitCode := m.Run()

	// Commit the transaction after tests are completed
	if err := tx.Commit(context.Background()); err != nil {
		log.Fatal("cannot commit transaction:", err)
	}

	// Close the connection
	if err := conn.Close(context.Background()); err != nil {
		log.Fatal("cannot close connection:", err)
	}

	// Exit with the test result code
	os.Exit(exitCode)
}

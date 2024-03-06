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
	defer tx.Rollback(context.Background())

	testQueries = New(tx)
	os.Exit(m.Run())
}

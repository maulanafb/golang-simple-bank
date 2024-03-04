// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64            `json:"id"`
	Owner     string           `json:"owner"`
	Balance   int64            `json:"balance"`
	Currency  string           `json:"currency"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be positive or negative
	Amount    int64            `json:"amount"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must positive
	Amount    int64            `json:"amount"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
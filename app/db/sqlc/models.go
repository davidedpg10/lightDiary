// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Entry struct {
	ID        int32        `json:"id"`
	Title     string       `json:"title"`
	Message   string       `json:"message"`
	Mood      []string     `json:"mood"`
	CreatedAt sql.NullTime `json:"created_at"`
}

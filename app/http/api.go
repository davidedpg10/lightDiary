package api

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/davidedpg10/golang-diary/db/sqlc"
)

func Run(queries *db.Queries) {
	http.HandleFunc("/api/entry/list", ListEntries(queries))
	http.ListenAndServe(":8080", nil)
}

func ListEntries(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entries, err := q.ListEntries(context.Background())
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}
		w.Write([]byte(fmt.Sprintf("%+v", entries)))

	}
}

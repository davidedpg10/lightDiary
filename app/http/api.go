package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/davidedpg10/lightDiary/db/sqlc"
)

type SingleEntryParams struct {
	ID int32 `json:"id"`
}

func Run(queries *db.Queries) {
	http.HandleFunc("/api/entry/get", GetEntry(queries))
	http.HandleFunc("/api/entry/list", ListEntries(queries))
	http.HandleFunc("/api/entry/add", AddEntry(queries))
	http.HandleFunc("/api/entry/update", UpdateEntry(queries))
	http.HandleFunc("/api/entry/delete", DeleteEntry(queries))
	http.ListenAndServe(":8080", nil)
}

func GetEntry(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validating and converting parameter
		strId := r.URL.Query().Get("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: Invalid, or missing parameter 'id'"))
			return
		}

		// Getting the entry
		entry, err := q.GetEntry(context.Background(), int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// converting the entry to JSON
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}

		// Returning the entry
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", string(entryJSON))))
	}
}

func AddEntry(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body
		decoder := json.NewDecoder(r.Body)
		var params db.CreateEntryParams
		err := decoder.Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// Inserting the entry
		entry, err := q.CreateEntry(context.Background(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// converting the entry to JSON
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}

		// Returning the new entry
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", string(entryJSON))))
	}
}

func ListEntries(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting the entries
		entries, err := q.ListEntries(context.Background())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}

		// Converting the entries to JSON
		entriesJSON, err := json.Marshal(entries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}

		// Returning the entries
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", string(entriesJSON))))

	}
}

func UpdateEntry(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body
		decoder := json.NewDecoder(r.Body)
		var params db.UpdateEntryParams
		err := decoder.Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// Updating the entry
		entry, err := q.UpdateEntry(context.Background(), params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// converting the entry to JSON
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		}

		// Returning the new entry
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", string(entryJSON))))
	}
}

func DeleteEntry(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decoding the request body
		decoder := json.NewDecoder(r.Body)
		var params SingleEntryParams
		err := decoder.Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// Deleting the entry
		_, err = q.DeleteEntry(context.Background(), params.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}

		// Returning the deleted entry
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Entry deleted"))
	}
}

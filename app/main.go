package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/davidedpg10/golang-diary/db/sqlc"
	api "github.com/davidedpg10/golang-diary/http"
	_ "github.com/lib/pq"
)

var dbHost, dbUser, dbPass, dbName, dbDriver, dbSource string // postgresql conn details
var dbPort int                                                // postgresql port

func init() {
	dbHost = os.Getenv("POSTGRES_HOST")
	dbPort = 5432
	dbUser = os.Getenv("POSTGRES_USER")
	dbPass = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbDriver = "postgres"
	dbSource = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
}

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	queries := db.New(conn)

	api.Run(queries)

}

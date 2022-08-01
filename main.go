package main

import (
	"database/sql"
	"log"

	"github.com/dhiyaaulauliyaa/learn-go/api"
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/kajian?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error when connecting to databse: ", err)
	}

	server, err := api.NewServer(db.NewStore(conn))
	if err != nil {
		log.Fatal("Error when creating server: ", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error when starting server: ", err)
	}

}

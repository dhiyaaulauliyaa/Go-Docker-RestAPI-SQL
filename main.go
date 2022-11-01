package main

import (
	"database/sql"
	"log"

	"github.com/dhiyaaulauliyaa/learn-go/api"
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/dhiyaaulauliyaa/learn-go/util"
	_ "github.com/lib/pq"
)

func main() {
	/* Load Config File */
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}

	/* Connect Database */
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error when connecting to databse: ", err)
	}

	/* Create Server */
	server, err := api.NewServer(db.NewStore(conn))
	if err != nil {
		log.Fatal("Error when creating server: ", err)
	}

	/* Start Server */
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error when starting server: ", err)
	}

}

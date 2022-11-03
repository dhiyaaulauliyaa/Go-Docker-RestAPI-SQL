package main

import (
	"database/sql"
	"log"

	"github.com/dhiyaaulauliyaa/learn-go/api"
	"github.com/dhiyaaulauliyaa/learn-go/config"
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/dhiyaaulauliyaa/learn-go/util"
	_ "github.com/lib/pq"
)

func main() {
	/* Load Config File */
	appConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error when loading config: ", err)
	}

	/* Connect Database */
	conn, err := sql.Open(appConfig.DBDriver, appConfig.DBSource)
	if err != nil {
		log.Fatal("Error when connecting to databse: ", err)
	}

	/* Init Firebase */
	firebaseApp := config.SetupFirebase()

	/* Create Server */
	server, err := api.NewServer(
		db.NewStore(conn),
		*config.SetupFirebaseAuth(firebaseApp),
	)
	if err != nil {
		log.Fatal("Error when creating server: ", err)
	}

	/* Start Server */
	err = server.Start(appConfig.ServerAddress)
	if err != nil {
		log.Fatal("Error when starting server: ", err)
	}

}

package main

import (
	"database/sql"
	"ecom/api"
	"ecom/db/sqlc"
	"ecom/db/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatal("error connection to DB", err)
	}
	store := sqlc.NewStore(db)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("error while craeting new server", err)
	}
	err = server.Start(config.HttpServerAddr)
	if err != nil {
		log.Fatal("error while running the server", err)
	}
}

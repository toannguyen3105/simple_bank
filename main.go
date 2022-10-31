package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/toannguyen3105/simplebank/api"
	db "github.com/toannguyen3105/simplebank/db/sqlc"
	"github.com/toannguyen3105/simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

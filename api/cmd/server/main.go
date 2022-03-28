package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
	"github.com/heritechie/motiket/api/internal/transport/http"
	"github.com/heritechie/motiket/api/internal/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	fmt.Println(config.DBSource)

	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := http.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

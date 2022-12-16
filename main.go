package main

import (
	"at01/api"
	db "at01/db/sqlc"
	"at01/initializers"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://pgVanThinh:VanThinh2512.@localhost:8080/pt01_StoreMobile?sslmode=disable"
	serverAddress = "http://192.168.1.19:8888/"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	// initializers.SyncDatabase()

}

func main() {

	store := db.NewStore(initializers.DB)
	server := api.NewServer(store)

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

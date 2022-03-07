package main

import (
	"github.com/FSDelite/go-rest-api/database"
	"github.com/FSDelite/go-rest-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}

package main

import "github.com/FSDelite/go-rest-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}

package main

import "api-v1/src/server"

func main() {
	srv := server.NewServer("localhost", "8080")
	srv.Run()

}
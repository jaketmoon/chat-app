package main

import (
	"chat-app/internal/server"
	"fmt"
)

func main() {
	server := server.NewServer()
	err := server.ListenAndServer()
	if err != nil {
		panic(fmt.Sprintf("cannot start server", err))
	}
}

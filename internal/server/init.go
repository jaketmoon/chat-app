package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   *mongo.Database
	ws   *websocket.WebsocketServer
}

func NewServer() *http.Server {
	portStr := os.Getenv("PORT") //
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080
	}
	db, err := database.New()
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}
	conversationService := service.NewConversationService(db)
	messageService := service.NewMessageService(db)
	ws := websocket.NewWebsocketServer(conversationService, messageService)
	go ws.Run()
	newServer := &Server{
		port: port,
		db:   db,
		ws:   ws,
	}
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}

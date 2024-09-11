package server

import (
	"chat-app/internal/database"
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
	portStr := os.Getenv("PORT")       //获取环境变量PORT
	port, err := strconv.Atoi(portStr) //让portStr转换为int类型
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
	go ws.Run() //启动websocket服务
	newServer := &Server{
		port: port,
		db:   db,
		ws:   ws,
	}
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRouter(),
		IdleTimeout:  time.Minute,      //空闲超时时间
		ReadTimeout:  10 * time.Second, //读取超时时间
		WriteTimeout: 30 * time.Second, //写入超时时间
	}
	return server
}

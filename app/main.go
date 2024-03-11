package main

import (
	"github.com/gin-gonic/gin"
	"go-api-server/chatroom/handler"
	"go-api-server/chatroom/repository/mysql"
	"go-api-server/chatroom/usecase"
	"log"
)

func main() {
	server := gin.Default()
	db, err := mysql.NewMysqlRepository()
	if err != nil {
		panic(err)
	}

	chatroomUseCase := usecase.NewChatroomUseCase(db)
	handler.NewChatroomHandler(server, chatroomUseCase)
	log.Fatal(server.Run(":8080"))
}

package main

import (
	"github.com/gin-gonic/gin"
	"go-api-server/internal/chatroom/handler"
	"go-api-server/internal/chatroom/repository"
	"go-api-server/internal/chatroom/usecase"
	"go-api-server/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	server := gin.Default()

	// init
	dsn := "id:pw@tcp(127.0.0.1:3306)/hama?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(domain.Chatroom{})

	chatroomUseCase := usecase.NewChatroomUseCase(repository.NewMysqlRepository(db))
	handler.NewChatroomHandler(server, chatroomUseCase)
	log.Fatal(server.Run(":8080"))
}

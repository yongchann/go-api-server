package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-server/api/handler"
	"go-api-server/config"
	"go-api-server/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	server := gin.Default()

	db := initDB()

	handler.NewChatroomRouter(server, db)
	handler.NewMemberRouter(server, db)

	log.Fatal(server.Run(":8080"))
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/hama?charset=utf8mb4&parseTime=True&loc=Local", config.Get().Db.User, config.Get().Db.Passwrod)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	_ = db.AutoMigrate(entity.Chatroom{})
	_ = db.AutoMigrate(entity.Member{})
	return db
}

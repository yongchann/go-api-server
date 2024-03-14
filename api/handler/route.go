package handler

import (
	"github.com/gin-gonic/gin"
	"go-api-server/internal/domain/chatroom"
	"go-api-server/internal/domain/member"
	"gorm.io/gorm"
)

func NewChatroomRouter(rg *gin.Engine, db *gorm.DB) {
	repo := chatroom.NewRepository(db)
	usecase := chatroom.NewUseCase(repo)
	handler := ChatroomHandler{chatroomUseCase: usecase}

	rg.POST("/chatrooms", handler.CreateChatroom)
	rg.GET("/chatrooms", handler.FetchChatroom)
	rg.GET("/chatrooms/:id", handler.GetChatroom)
	rg.PUT("/chatrooms/:id", handler.UpdateChatroom)
	rg.DELETE("/chatrooms/:id", handler.DeleteChatroom)
}

func NewMemberRouter(rg *gin.Engine, db *gorm.DB) {
	repo := member.NewRepository(db)
	usecase := member.NewUseCase(repo)
	handler := MemberHandler{memberUseCase: usecase}

	rg.POST("/user", handler.Join)
	rg.GET("/user/:nickname", handler.FindByNickname)
}

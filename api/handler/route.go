package handler

import (
	"github.com/gin-gonic/gin"
	"go-api-server/api/middleware"
	"go-api-server/internal/domain/chatroom"
	"go-api-server/internal/domain/member"
	"gorm.io/gorm"
)

func NewChatroomRouter(rg *gin.Engine, db *gorm.DB) {
	repo := chatroom.NewRepository(db)
	usecase := chatroom.NewUseCase(repo)
	handler := ChatroomHandler{chatroomUseCase: usecase}

	chatroomRg := rg.Group("/chatrooms").Use(middleware.JwtAuthMiddleware())
	{
		chatroomRg.POST("", handler.CreateChatroom)
		chatroomRg.GET("", handler.FetchChatroom)
		chatroomRg.GET("/:id", handler.GetChatroom)
		chatroomRg.PUT("/:id", handler.UpdateChatroom)
		chatroomRg.DELETE("/:id", handler.DeleteChatroom)
	}
}

func NewMemberRouter(rg *gin.Engine, db *gorm.DB) {
	repo := member.NewRepository(db)
	usecase := member.NewUseCase(repo)
	handler := MemberHandler{memberUseCase: usecase}

	memberRg := rg.Group("/members")
	{
		memberRg.POST("/", handler.Join)
		memberRg.GET("/:nickname", handler.FindByNickname)
		memberRg.POST("/login", handler.Login)
	}
}

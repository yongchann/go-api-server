package handler

import (
	"github.com/gin-gonic/gin"
	"go-api-server/domain"
)

type ChatroomHandler struct {
	chatroomUseCase domain.ChatroomUseCase
}

func NewChatroomHandler(rg *gin.Engine, us domain.ChatroomUseCase) {
	handler := ChatroomHandler{
		chatroomUseCase: us,
	}

	rg.GET("/chatrooms", handler.FetchChatroom)
	rg.GET("/chatrooms/{id}", handler.GetChatroom)
	rg.POST("/chatrooms", handler.CreateChatroom)
	rg.PUT("/chatrooms", handler.UpdateChatroom)
	rg.DELETE("/chatrooms/{id}", handler.DeleteChatroom)
}

func (cr *ChatroomHandler) FetchChatroom(context *gin.Context) {

}

func (cr *ChatroomHandler) GetChatroom(context *gin.Context) {

}

func (cr *ChatroomHandler) DeleteChatroom(context *gin.Context) {

}

func (cr *ChatroomHandler) CreateChatroom(context *gin.Context) {

}

func (cr *ChatroomHandler) UpdateChatroom(context *gin.Context) {

}

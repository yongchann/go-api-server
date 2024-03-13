package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-api-server/internal/chatroom/repository"
	"go-api-server/internal/domain"
	"net/http"
	"strconv"
)

type ChatroomHandler struct {
	chatroomUseCase domain.ChatroomUseCase
}

func NewChatroomHandler(rg *gin.Engine, us domain.ChatroomUseCase) {
	handler := ChatroomHandler{
		chatroomUseCase: us,
	}

	rg.POST("/chatrooms", handler.CreateChatroom)
	rg.GET("/chatrooms", handler.FetchChatroom)
	rg.GET("/chatrooms/:id", handler.GetChatroom)
	rg.PUT("/chatrooms/:id", handler.UpdateChatroom)
	rg.DELETE("/chatrooms/:id", handler.DeleteChatroom)
}

func (cr *ChatroomHandler) CreateChatroom(c *gin.Context) {
	var newChatroom domain.Chatroom
	if err := c.BindJSON(&newChatroom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := cr.chatroomUseCase.Create(&newChatroom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newChatroom)
}

func (cr *ChatroomHandler) FetchChatroom(c *gin.Context) {
	chatrooms, err := cr.chatroomUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chatrooms)
}

func (cr *ChatroomHandler) GetChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	chatroom, err := cr.chatroomUseCase.FindById(int64(pId))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chatroom)
}

func (cr *ChatroomHandler) UpdateChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var updated domain.Chatroom
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := cr.chatroomUseCase.Update(int64(pId), &updated); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (cr *ChatroomHandler) DeleteChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := cr.chatroomUseCase.Delete(int64(pId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

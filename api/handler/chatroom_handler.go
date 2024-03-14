package handler

import (
	"github.com/gin-gonic/gin"
	"go-api-server/internal/entity"
	"net/http"
	"strconv"
)

type ChatroomHandler struct {
	chatroomUseCase entity.ChatroomUseCase
}

func (h *ChatroomHandler) CreateChatroom(c *gin.Context) {
	var newChatroom entity.Chatroom
	if err := c.BindJSON(&newChatroom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.chatroomUseCase.Create(&newChatroom); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newChatroom)
}

func (h *ChatroomHandler) FetchChatroom(c *gin.Context) {
	chatrooms, err := h.chatroomUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chatrooms)
}

func (h *ChatroomHandler) GetChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	chatroom, err := h.chatroomUseCase.FindById(int64(pId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chatroom)
}

func (h *ChatroomHandler) UpdateChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var updated entity.Chatroom
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.chatroomUseCase.Update(int64(pId), &updated); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ChatroomHandler) DeleteChatroom(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.chatroomUseCase.Delete(int64(pId)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

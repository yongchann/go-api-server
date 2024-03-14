package handler

import (
	"github.com/gin-gonic/gin"
	"go-api-server/internal/entity"
	"net/http"
)

type MemberHandler struct {
	memberUseCase entity.MemberUseCase
}

func (h *MemberHandler) Join(c *gin.Context) {
	var newMember entity.Member
	if err := c.BindJSON(&newMember); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.memberUseCase.Join(&newMember); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newMember)
}

func (h *MemberHandler) FindByNickname(c *gin.Context) {
	pNickname := c.Param("nickname")
	findMember, err := h.memberUseCase.FindByNickname(pNickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, findMember)
}

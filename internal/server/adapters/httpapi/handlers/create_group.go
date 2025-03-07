package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) CreateGroup(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	var request createGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, newStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	groupId, err := h.model.CreateGroup(
		c,
		actorId,
		request.Name,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, createGroupResponse{Id: groupId, Status: status{Code: http.StatusOK}})
}

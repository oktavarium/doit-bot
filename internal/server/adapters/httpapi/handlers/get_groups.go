package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) GetGroups(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	groups, err := h.model.GetGroups(c, actorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, getGroupsResponse{Groups: groups, Status: status{Code: http.StatusOK}})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) DeleteTaskById(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	var request deleteTaskByIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.model.DeleteTaskById(
		c,
		actorId,
		request.Id,
	); err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewStatusResponse(http.StatusOK, ""))
}

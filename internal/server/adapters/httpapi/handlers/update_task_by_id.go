package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) UpdateTaskById(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	var request updateTaskByIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.model.UpdateTaskById(
		c,
		actorId,
		request.Id,
		request.AssigneeId,
		request.ListId,
		request.Name,
		request.Description,
		request.Done,
	); err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.NewStatusResponse(http.StatusOK, ""))
}

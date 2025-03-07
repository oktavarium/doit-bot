package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) GetTaskById(c *gin.Context) {
	var request getTaskByIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	task, err := h.model.GetTaskById(
		c,
		request.Id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, getTaskByIdResponse{Task: task, Status: common.Status{Code: http.StatusOK}})
}

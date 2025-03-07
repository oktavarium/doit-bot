package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetTaskById(c *gin.Context) {
	var request getTaskByIdRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, newStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	task, err := h.model.GetTaskById(
		c,
		request.Id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, getTaskByIdResponse{Task: task, Status: status{Code: http.StatusOK}})
}

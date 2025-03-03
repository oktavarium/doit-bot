package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetTaskById(c *gin.Context) {
	var request getTaskByIdRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateGetTaskByIdRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.model.GetTaskById(
		c,
		request.Id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getTaskResponse{task})
}

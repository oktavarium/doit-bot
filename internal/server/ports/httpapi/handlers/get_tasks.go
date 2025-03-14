package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (h *Handlers) GetTasks(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	cmd := query.GetTasks{
		UserId: actorId,
	}

	tasks, err := h.app.Queries.GetTasks.Handle(c, cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	result := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, Task{
			Id:          task.Id(),
			OwnerId:     task.OwnerId(),
			Name:        task.Name(),
			Description: task.Description(),
			Status:      task.Status(),
		})
	}

	c.JSON(http.StatusOK, TasksResponse{Tasks: result, Status: newStatusResponse(http.StatusOK, "")})
}

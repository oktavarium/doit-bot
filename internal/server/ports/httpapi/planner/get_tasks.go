package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) GetTasks(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	cmd := query.GetTasks{
		UserId: actorId,
	}

	tasks, err := p.app.Queries.GetTasks.Handle(c, cmd)
	if err != nil {
		common.ErrorToContext(c, common.FromAppError(err))
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

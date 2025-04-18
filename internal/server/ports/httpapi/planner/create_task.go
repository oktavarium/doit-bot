package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) CreateTask(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	var request NewTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorToContext(c, common.NewBadRequestError(err))
		return
	}

	cmd := command.CreateTask{
		OwnerId:     actorId,
		ListId:      request.ListId,
		Name:        request.Name,
		Description: request.Description,
	}

	taskId, err := p.app.Commands.CreateTask.Handle(c, cmd)
	if err != nil {
		common.ErrorToContext(c, common.FromAppError(err))
		return
	}

	c.JSON(http.StatusCreated, TaskIdResponse{Id: taskId, Status: newStatusResponse(http.StatusCreated, "")})
}

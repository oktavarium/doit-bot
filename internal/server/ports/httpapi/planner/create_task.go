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
		Name:        request.Name,
		Description: request.Description,
	}

	if err := p.app.Commands.CreateTask.Handle(c, cmd); err != nil {
		common.ErrorToContext(c, common.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, TaskIdResponse{Id: "", Status: newStatusResponse(http.StatusOK, "")})
}

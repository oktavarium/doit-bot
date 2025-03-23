package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) CreateList(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	var request NewListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorToContext(c, common.NewBadRequestError(err))
		return
	}

	cmd := command.CreateList{
		OwnerId:     actorId,
		Name:        request.Name,
		Description: request.Description,
	}

	taskId, err := p.app.Commands.CreateList.Handle(c, cmd)
	if err != nil {
		common.ErrorToContext(c, common.FromAppError(err))
		return
	}

	c.JSON(http.StatusCreated, ListIdResponse{Id: taskId, Status: newStatusResponse(http.StatusCreated, "")})
}

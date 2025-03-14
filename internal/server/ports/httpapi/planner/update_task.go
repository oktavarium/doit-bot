package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) UpdateTask(c *gin.Context, id string) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	var request UpdateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorToContext(c, common.NewBadRequestError(err))
		return
	}

	if request.Status != nil {
		cmd := command.SetTaskStatus{
			ActorId: actorId,
			TaskId:  id,
			Status:  *request.Status,
		}

		if err := p.app.Commands.SetTaskStatus.Handle(c, cmd); err != nil {
			common.ErrorToContext(c, common.NewInternalServerError(err))
			return
		}

		c.JSON(http.StatusOK, TaskIdResponse{Id: "", Status: newStatusResponse(http.StatusOK, "")})
	}
}

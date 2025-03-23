package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/command"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) UpdateList(c *gin.Context, id string) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	var request UpdateListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		common.ErrorToContext(c, common.NewBadRequestError(err))
		return
	}
	cmd := command.UpdateList{
		ActorId:     actorId,
		ListId:      id,
		Name:        request.Name,
		Description: request.Description,
	}

	if err := p.app.Commands.UpdateList.Handle(c, cmd); err != nil {
		common.ErrorToContext(c, common.FromAppError(err))
		return
	}

	c.JSON(http.StatusOK, newStatusResponse(http.StatusOK, ""))

}

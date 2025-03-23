package planner

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/app/query"
	"github.com/oktavarium/doit-bot/internal/server/ports/httpapi/common"
)

func (p *Planner) GetLists(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		common.ErrorToContext(c, common.NewInternalServerError(errors.New("empty context")))
		return
	}

	cmd := query.GetLists{
		UserId: actorId,
	}

	tasks, err := p.app.Queries.GetLists.Handle(c, cmd)
	if err != nil {
		common.ErrorToContext(c, common.FromAppError(err))
		return
	}

	result := make([]List, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, List{
			Id:          task.Id(),
			OwnerId:     task.OwnerId(),
			Name:        task.Name(),
			Description: task.Description(),
		})
	}

	c.JSON(http.StatusOK, ListsResponse{Lists: result, Status: newStatusResponse(http.StatusOK, "")})
}

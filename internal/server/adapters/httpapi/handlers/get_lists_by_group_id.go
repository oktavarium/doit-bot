package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
)

func (h *Handlers) GetListsByGroupId(c *gin.Context) {
	actorId, ok := common.ActorIdFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
		return
	}

	var request getListsByGroupIdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, common.NewStatusResponse(http.StatusBadRequest, err.Error()))
		return
	}

	lists, err := h.model.GetListsByGroupId(
		c,
		actorId,
		request.Id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, getListsByGroupIdResponse{Lists: lists, Status: common.Status{Code: http.StatusOK}})
}

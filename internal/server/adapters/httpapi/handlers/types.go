package handlers

import (
	"github.com/oktavarium/doit-bot/internal/server/adapters/httpapi/common"
	"github.com/oktavarium/doit-bot/internal/server/dto"
)

type createTaskRequest struct {
	AssigneeId  *string `json:"assignee_id,omitempty" validate:"omitnil,required,eq=24"`
	ListId      *string `json:"list_id,omitempty" validate:"omitnil,required,eq=24"`
	Name        string  `json:"name" validate:"required,lte=32"`
	Description string  `json:"description,omitempty" validate:"omitempty,required,lte=256"`
}

type createTaskResponse struct {
	Id     string        `json:"id"`
	Status common.Status `json:"status"`
}

type createGroupRequest struct {
	Name string `json:"name,omitempty" validate:"required,lte=32"`
}

type createGroupResponse struct {
	Id     string        `json:"id"`
	Status common.Status `json:"status"`
}

type getGroupsResponse struct {
	Groups []*dto.Group  `json:"groups"`
	Status common.Status `json:"status"`
}

type createListRequest struct {
	GroupId     *string `json:"group_id,omitempty" validate:"omitnil,required,eq=24"`
	Name        string  `json:"name" validate:"required,lte=32"`
	Description string  `json:"description,omitempty" validate:"omitempty,required,lte=256"`
}

type createListResponse struct {
	Id     string        `json:"id"`
	Status common.Status `json:"status"`
}

type deleteTaskByIdRequest struct {
	Id string `json:"id" validate:"required,eq=24"`
}

type getTaskByIdRequest struct {
	Id string `json:"id" validate:"required,eq=24"`
}

type getTaskByIdResponse struct {
	*dto.Task
	Status common.Status `json:"status"`
}

type getListsByGroupIdRequest struct {
	Id string `json:"id" validate:"required,eq=24"`
}

type getListsByGroupIdResponse struct {
	Lists  []*dto.List   `json:"lists"`
	Status common.Status `json:"status"`
}

type getTasksByOwnerResponse struct {
	Tasks  []*dto.Task   `json:"tasks"`
	Status common.Status `json:"status"`
}

type getTasksResponse struct {
	Tasks  []*dto.Task   `json:"tasks"`
	Status common.Status `json:"status"`
}

type updateTaskByIdRequest struct {
	Id          string  `json:"id" validate:"required,eq=24"`
	AssigneeId  *string `json:"assignee_id,omitempty" validate:"omitnil,required,eq=24"`
	ListId      *string `json:"list_id,omitempty" validate:"omitnil,required,eq=24"`
	Name        *string `json:"name,omitempty" validate:"omitnil,required,lte=32"`
	Description *string `json:"description,omitempty" validate:"omitnil,required,lte=256"`
	Done        *bool   `json:"done,omitempty" validate:"omitnil"`
}

type setTaskDoneByIdRequest struct {
	Id   string `json:"id" validate:"required,eq=24"`
	Done bool   `json:"done" validate:"required"`
}

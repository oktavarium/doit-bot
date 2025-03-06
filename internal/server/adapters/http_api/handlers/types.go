package handlers

import "github.com/oktavarium/doit-bot/internal/server/dto"

type createTaskRequest struct {
	AssigneeId  *string `json:"assignee_id,omitempty" validation:"omitnil,required,lte=32"`
	ListId      *string `json:"list_id,omitempty" validation:"required"`
	Name        string  `json:"name" validation:"required,lte=32"`
	Description string  `json:"description,omitempty" validation:"omitempty,required,lte=256"`
}

type createTaskResponse struct {
	Id string `json:"id" validation:"required"`
}

type createGroupRequest struct {
	Name string `json:"name,omitempty" validation:"omitempty,required,lte=32"`
}

type createGroupResponse struct {
	Id string `json:"id" validation:"required"`
}

type getGroupsResponse struct {
	Groups []*dto.Group `json:"groups"`
}

type createListRequest struct {
	GroupId     *string `json:"group_id,omitempty" validation:"omitnil,required"`
	Name        string  `json:"name" validation:"required,lte=32"`
	Description string  `json:"description,omitempty" validation:"omitempty,required,lte=256"`
}

type createListResponse struct {
	Id string `json:"id" validation:"required"`
}

type deleteTaskByIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getTaskByIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getTaskByIdResponse struct {
	*dto.Task
}

type getListsByGroupIdRequest struct {
	GroupId string `json:"group_id" validation:"required"`
}

type getListsByGroupIdResponse struct {
	Lists []*dto.List `json:"lists"`
}

type getTasksByOwnerResponse struct {
	Tasks []*dto.Task `json:"tasks"`
}

type updateTaskByIdRequest struct {
	Id          string  `json:"id" validation:"required"`
	AssigneeId  *string `json:"assignee_id,omitempty" validation:"omitnil,required,lte=32"`
	ListId      *string `json:"list_id,omitempty" validation:"omitnil,required,lte=32"`
	Name        *string `json:"name,omitempty" validation:"omitnil,required,lte=32"`
	Description *string `json:"description,omitempty" validation:"omitnil,required,lte=256"`
	Done        *bool   `json:"done,omitempty" validation:"omitnil"`
}

type setTaskDoneByIdRequest struct {
	Id   string `json:"id" validation:"required"`
	Done bool   `json:"done" validation:"required"`
}

package handlers

import "github.com/oktavarium/doit-bot/internal/server/dto"

type status struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type statusResponse struct {
	Status status `json:"status"`
}

func newStatusResponse(code int64, message string) statusResponse {
	return statusResponse{
		Status: status{
			Code:    code,
			Message: message,
		},
	}
}

type createTaskRequest struct {
	AssigneeId  *string `json:"assignee_id,omitempty" validation:"omitnil,required,lte=32"`
	ListId      *string `json:"list_id,omitempty" validation:"required"`
	Name        string  `json:"name" validation:"required,lte=32"`
	Description string  `json:"description,omitempty" validation:"omitempty,required,lte=256"`
}

type createTaskResponse struct {
	Id     string `json:"id" validation:"required"`
	Status status `json:"status"`
}

type createGroupRequest struct {
	Name string `json:"name,omitempty" validation:"omitempty,required,lte=32"`
}

type createGroupResponse struct {
	Id     string `json:"id" validation:"required"`
	Status status `json:"status"`
}

type getGroupsResponse struct {
	Groups []*dto.Group `json:"groups"`
	Status status       `json:"status"`
}

type createListRequest struct {
	GroupId     *string `json:"group_id,omitempty" validation:"omitnil,required"`
	Name        string  `json:"name" validation:"required,lte=32"`
	Description string  `json:"description,omitempty" validation:"omitempty,required,lte=256"`
}

type createListResponse struct {
	Id     string `json:"id" validation:"required"`
	Status status `json:"status"`
}

type deleteTaskByIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getTaskByIdRequest struct {
	Id     string `json:"id" validation:"required"`
	Status status `json:"status"`
}

type getTaskByIdResponse struct {
	*dto.Task
	Status status `json:"status"`
}

type getListsByGroupIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getListsByGroupIdResponse struct {
	Lists  []*dto.List `json:"lists"`
	Status status      `json:"status"`
}

type getTasksByOwnerResponse struct {
	Tasks  []*dto.Task `json:"tasks"`
	Status status      `json:"status"`
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

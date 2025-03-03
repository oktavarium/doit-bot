package handlers

import "github.com/oktavarium/doit-bot/internal/server/internal/dto"

type createTaskRequest struct {
	Assignee    *string `json:"assignee_id,omitempty" validation:"omitnil,required,lte=32"`
	List        *string `json:"list_id,omitempty" validation:"omitnil,required,lte=32"`
	Summary     string  `json:"summary,omitempty" validation:"omitnil,required,lte=32"`
	Description *string `json:"description,omitempty" validation:"omitnil,required,lte=256"`
}

type createTaskResponse struct {
	Id string `json:"id" validation:"required"`
}

type deleteTaskByIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getTaskByIdRequest struct {
	Id string `json:"id" validation:"required"`
}

type getTaskResponse struct {
	*dto.Task
}

type getTasksByOwnerResponse struct {
	Tasks []*dto.Task `json:"tasks"`
}

type updateTaskByIdRequest struct {
	Id          string  `json:"id" validation:"required"`
	Assignee    *string `json:"assignee_id,omitempty" validation:"omitnil,required,lte=32"`
	List        *string `json:"list_id,omitempty" validation:"omitnil,required,lte=32"`
	Summary     *string `json:"summary,omitempty" validation:"omitnil,required,lte=32"`
	Description *string `json:"description,omitempty" validation:"omitnil,required,lte=256"`
	Done        *bool   `json:"done,omitempty" validation:"omitnil"`
}

type setTaskDoneByIdRequest struct {
	Id   string `json:"id" validation:"required"`
	Done bool   `json:"done" validation:"required"`
}

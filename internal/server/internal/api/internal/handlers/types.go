package handlers

import "github.com/oktavarium/doit-bot/internal/server/internal/dto"

type createTaskRequest struct {
	Assignee *int64 `json:"assignee,omitempty"`
	Summary  string `json:"summary,omitempty"`
}

type createTaskResponse struct {
	Id string `json:"id"`
}

type deleteTaskRequest struct {
	Id string `json:"id,omitempty"`
}

type getTaskRequest struct {
	Id string `json:"id,omitempty"`
}

type getTaskResponse struct {
	*dto.Task
}

type getTasksResponse struct {
	Tasks []*dto.Task `json:"tasks"`
}

type updateTaskRequest struct {
	Id       string  `json:"id,omitempty"`
	Assignee *int64  `json:"assignee,omitempty"`
	Summary  *string `json:"summary,omitempty"`
	Done     *bool   `json:"done,omitempty"`
}

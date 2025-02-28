package dto

type Task struct {
	Id       string `json:"id"`
	Owner    int64  `json:"owner"`
	Assignee *int64  `json:"assignee,omitempty"`
	Summary  string `json:"summary"`
	Done     bool   `json:"done"`
}

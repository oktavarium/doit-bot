package dto

type Task struct {
	Id          *string `json:"id,omitempty"`
	Owner       *string `json:"owner,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	Done        *bool   `json:"done,omitempty"`
}

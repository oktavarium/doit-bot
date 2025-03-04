package dto

type Group struct {
	Id      string `json:"id"`
	TgId    int64  `json:"tg_id,omitempty"`
	OwnerId string `json:"owner_id"`
	Name    string `json:"name"`
}

type List struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	GroupId     string `json:"group_id"`
	Name        string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

type Task struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	AssigneeId  string `json:"assignee_id,omitempty"`
	ListId      string `json:"list_id,omitempty"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type User struct {
	Id        string `json:"id"`
	TgId      int64  `json:"tg_id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

package dto

type Task struct {
	Id          string `json:"id"`
	Owner       string `json:"owner_id"`
	Assignee    string `json:"assignee_id,omitempty"`
	List        string `json:"list_id,omitempty"`
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

type List struct {
	Id          string `json:"id"`
	Owner       string `json:"owner_id"`
	Group       string `json:"group_id"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

type Group struct {
	Id          string `json:"id"`
	TgId        int64  `json:"tg_id,omitempty"`
	Owner       string `json:"owner_id"`
	Group       string `json:"last_name"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

package dto

type GroupType int64

const (
	WithChat = iota
	WithoutChat
)

type Group struct {
	Id      string    `json:"id"`
	OwnerId string    `json:"owner_id,omitempty"`
	TgId    int64     `json:"tg_id,omitempty"`
	Type    GroupType `json:"type"`
	Users   []string  `json:"users,omitempty"`
	Name    string    `json:"name"`
}

type List struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	GroupId     string `json:"group_id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Task struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	AssigneeId  string `json:"assignee_id,omitempty"`
	ListId      string `json:"list_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}

type User struct {
	Id        string   `json:"id"`
	TgId      int64    `json:"tg_id"`
	Groups    []string `json:"groups,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Username  string   `json:"username,omitempty"`
}

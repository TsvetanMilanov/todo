package models

import "time"

// UserRequestModel model.
type UserRequestModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequestModel model.
type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TodoRequestModel model.
type TodoRequestModel struct {
	ID       string     `json:"id"`
	Content  string     `json:"content"`
	Category string     `json:"category"`
	Deadline *time.Time `json:"deadline"`
	Priority int        `json:"priority"`
}

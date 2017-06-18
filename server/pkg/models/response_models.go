package models

// UserResponseModel model.
type UserResponseModel struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

package models

// UserRequestModel model.
type UserRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequestModel model.
type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

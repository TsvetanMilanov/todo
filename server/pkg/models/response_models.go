package models

// UserResponseModel model.
type UserResponseModel struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// TokenResponseModel model.
type TokenResponseModel struct {
	AccessToken string   `json:"access_token"`
	TokenType   string   `json:"token_type"`
	ExpiresIn   int      `json:"expires_in"`
	Scopes      []string `json:"scopes"`
}

// LoginResponseModel model.
type LoginResponseModel struct {
	TokenResponseModel
	User UserResponseModel `json:"user"`
}

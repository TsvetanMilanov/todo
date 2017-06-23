package types

// ServerConfig is the server configuration.
type ServerConfig struct {
	Host  string `json:"host"`
	Proto string `json:"proto"`
}

// LoginRequest username and password login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserData is the user information.
type UserData struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// LoginResponse is the response after successfull login.
type LoginResponse struct {
	AccessToken string   `json:"access_token"`
	TokenType   string   `json:"token_type"`
	ExpiresIn   uint64   `json:"expires_in"`
	Scopes      []string `json:"scopes"`
	User        UserData `json:"user"`
}

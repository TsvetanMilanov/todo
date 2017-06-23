package api

import "github.com/TsvetanMilanov/todo/clients/cli/pkg/types"

// AuthAPIService authentication api related methods.
type AuthAPIService struct {
	ServerClient types.IServerClient `inject:"serverClient"`
}

// Login sends login request to the api and returns the response.
func (auth *AuthAPIService) Login(username, password string) (*types.LoginResponse, error) {
	result := types.LoginResponse{}
	body := types.LoginRequest{Username: username, Password: password}
	err := auth.ServerClient.Post("auth/login", &body, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

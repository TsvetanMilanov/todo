package services

import "github.com/TsvetanMilanov/todo/clients/cli/pkg/types"

// AuthService authentication related methods.
type AuthService struct {
	AuthAPIService types.IAuthAPIService `inject:"authAPIService"`
	UserService    types.IUserService    `inject:"userService"`
}

// Login sends the username and password to the server and saves the response on successful login.
func (auth *AuthService) Login(username, password string) (*types.LoginResponse, error) {
	result, err := auth.AuthAPIService.Login(username, password)
	if err != nil {
		return nil, err
	}

	err = auth.UserService.SaveUser(*result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

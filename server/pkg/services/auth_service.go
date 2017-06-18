package services

import (
	"errors"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
)

// AuthService auth related methods
type AuthService struct {
	DbService     types.IDbService     `inject:"dbService"`
	UsersService  types.IUsersService  `inject:"usersService"`
	Helpers       types.IHelpers       `inject:"helpers"`
	TokensService types.ITokensService `inject:"tokensService"`
}

// Login creates new token for the user.
func (auth *AuthService) Login(username, password string) (*models.Token, error) {
	_, err := auth.AuthenticateUserWithPassword(username, password)
	if err != nil {
		return nil, err
	}

	token, err := auth.TokensService.CreateToken(username, constants.BearerTokenType)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// AuthenticateUserWithPassword checks if the username and the password are valid.
func (auth *AuthService) AuthenticateUserWithPassword(username, password string) (*models.User, error) {
	user, err := auth.UsersService.GetUser(username)
	if err != nil {
		return nil, err
	}

	err = auth.Helpers.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

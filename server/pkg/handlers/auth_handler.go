package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/labstack/echo"
)

// AuthHandler auth related methods.
type AuthHandler struct {
	AuthService  types.IAuthService  `inject:"authService"`
	UsersService types.IUsersService `inject:"usersService"`
}

// Login returns access token along with user info.
func (handler *AuthHandler) Login(context echo.Context) error {
	loginData := models.LoginRequestModel{}

	err := context.Bind(&loginData)

	if err != nil {
		return err
	}

	token, err := handler.AuthService.Login(loginData.Username, loginData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := handler.UsersService.GetUser(loginData.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tokenResponse := models.TokenResponseModel{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
		Scopes:      token.Scopes,
		TokenType:   token.TokenType,
	}

	userResponse := models.UserResponseModel{
		Username: user.Username,
		Roles:    user.Roles,
	}

	response := models.LoginResponseModel{
		TokenResponseModel: tokenResponse,
		User:               userResponse,
	}

	return context.JSON(http.StatusCreated, response)
}

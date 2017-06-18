package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/labstack/echo"
)

// UsersHandler users related operations.
type UsersHandler struct {
	UsersService types.IUsersService `inject:"usersService"`
}

// AddUser creates new user.
func (handler *UsersHandler) AddUser(context echo.Context) error {
	user := models.UserRequestModel{}

	err := context.Bind(&user)

	if err != nil {
		return err
	}

	result, err := handler.UsersService.AddUser(user.Username, user.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resultUser := models.UserResponseModel{
		Username: result.Username,
		Roles:    result.Roles,
	}

	return context.JSON(http.StatusCreated, resultUser)
}

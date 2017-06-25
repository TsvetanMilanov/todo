package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/labstack/echo"
)

// TodosHandler TODOs related operations.
type TodosHandler struct {
	TodosService types.ITodosService `inject:"todosService"`
	Helpers      types.IHelpers      `inject:"helpers"`
}

// AddTodo creates new TODO.
func (handler *TodosHandler) AddTodo(context echo.Context) error {
	todo := models.TodoRequestModel{}

	err := context.Bind(&todo)
	if err != nil {
		return err
	}

	user, err := handler.Helpers.GetUserFromContext(context)
	if err != nil {
		return err
	}

	savedTodo, err := handler.TodosService.AddTodo(todo.Content, todo.Category, user.Username, todo.Deadline, todo.Priority)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusCreated, savedTodo)
}

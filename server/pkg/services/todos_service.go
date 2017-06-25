package services

import (
	"time"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/google/uuid"
)

// TodoServices TODOs related methods.
type TodoServices struct {
	DbService      types.IDbService      `inject:"dbService"`
	ModelValidator types.IModelValidator `inject:"modelValidator"`
}

// AddTodo adds todo to the user.
func (todoService *TodoServices) AddTodo(content,
	category,
	userID string,
	deadline *time.Time,
	priority int) (*models.Todo, error) {
	err := todoService.ModelValidator.ValidateNewTodoData(content, userID)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	todo := &models.Todo{
		ID:       id.String(),
		UserID:   userID,
		Content:  content,
		Category: category,
		Priority: priority,
		Deadline: deadline,
		Done:     false,
	}

	todos := todoService.DbService.GetCollection(constants.TodosCollectionName)

	err = todos.Insert(todo)
	if err != nil {
		return nil, err
	}

	return todo, err
}

package services

import (
	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

// UsersService users related operations.
type UsersService struct {
	DbService      types.IDbService      `inject:"dbService"`
	Helpers        types.IHelpers        `inject:"helpers"`
	ModelValidator types.IModelValidator `inject:"modelValidator"`
}

//AddUser creates a new user in the database.
func (service *UsersService) AddUser(username, password string) (*models.User, error) {
	err := service.ModelValidator.ValidateUser(username, password)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	encryptedPass, err := service.Helpers.EncryptString(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       id.String(),
		Username: username,
		Password: encryptedPass,
		Roles:    []string{constants.UserRole},
		Todos:    []string{},
	}

	users := service.DbService.GetCollection(constants.UsersCollectionName)

	err = users.Insert(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UsersService) GetUser(username string) (*models.User, error) {
	users := service.DbService.GetCollection(constants.UsersCollectionName)
	user := models.User{}
	err := users.Find(bson.M{"username": username}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

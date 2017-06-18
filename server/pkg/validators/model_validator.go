package validators

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
)

const (
	usernameMinLength = 4
	usernameMaxLength = 50
	passwordMinLength = 8
	passwordMaxLength = 150
)

// ModelValidator validator for db models.
type ModelValidator struct {
	DbService types.IDbService `inject:"dbService"`
}

// ValidateUser validates the user information.
func (validator *ModelValidator) ValidateUser(username, password string) error {
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return fmt.Errorf("username should be between %d and %d symbols", usernameMinLength, usernameMaxLength)
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return fmt.Errorf("password should be between %d and %d symbols", passwordMinLength, passwordMaxLength)
	}

	users := validator.DbService.GetCollection(constants.UsersCollectionName)

	count, err := users.Find(bson.M{"username": username}).Count()

	if err != nil {
		return err
	}

	if count != 0 {
		return fmt.Errorf("user with username %s already exists", username)
	}

	return nil
}

// ValidateLoginData validates login data.
func (validator *ModelValidator) ValidateLoginData(username, password string) error {
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return fmt.Errorf("username should be between %d and %d symbols", usernameMinLength, usernameMaxLength)
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return fmt.Errorf("password should be between %d and %d symbols", passwordMinLength, passwordMaxLength)
	}

	return nil
}

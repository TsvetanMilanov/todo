package services

import (
	"path"

	"encoding/json"
	"io/ioutil"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/constants"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
)

// UserService user related methods.
type UserService struct {
	Helpers types.IHelpers `inject:"helpers"`
}

// SaveUser saves the login info of the current user.
func (userService *UserService) SaveUser(loginInfo types.LoginResponse) error {
	userDataFilePath, err := userService.getUserDataFilePath()
	if err != nil {
		return err
	}

	err = userService.Helpers.EnsureDirExists(path.Dir(userDataFilePath), 0644)
	if err != nil {
		return err
	}

	userData, err := json.MarshalIndent(loginInfo, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(userDataFilePath, userData, 0644)
	return err
}

func (userService *UserService) getUserDataFilePath() (string, error) {
	homeDir, err := userService.Helpers.GetCurrentUserHomeDir()
	if err != nil {
		return "", err
	}

	todoDataFolder := path.Join(homeDir, constants.TodoDataFolder)
	return path.Join(todoDataFolder, constants.UserDataFile), nil
}

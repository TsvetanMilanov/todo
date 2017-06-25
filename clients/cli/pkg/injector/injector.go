package injector

import (
	"fmt"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/services"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/services/api"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/services/api/config"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/util"
	"github.com/facebookgo/inject"
)

// CreateInjectorGraph registers all required server dependencies and returns their graph.
func CreateInjectorGraph() (*inject.Graph, error) {
	var (
		injector = new(inject.Graph)
		err      error
	)

	var helpers types.IHelpers = &util.Helpers{}
	var authService types.IAuthService = &services.AuthService{}
	var userService types.IUserService = &services.UserService{}
	var logger types.ILogger = &services.Logger{}
	var serverConfigManager types.IServerConfigManager = &config.ServerConfigManager{}
	var authAPIService types.IAuthAPIService = &api.AuthAPIService{}
	var serverClient types.IServerClient = &api.ServerClient{}

	err = injector.Provide(
		&inject.Object{Value: helpers, Name: "helpers"},
		&inject.Object{Value: authService, Name: "authService"},
		&inject.Object{Value: serverConfigManager, Name: "serverConfigManager"},
		&inject.Object{Value: authAPIService, Name: "authAPIService"},
		&inject.Object{Value: serverClient, Name: "serverClient"},
		&inject.Object{Value: userService, Name: "userService"},
		&inject.Object{Value: logger, Name: "logger"},
	)

	if err != nil {
		return nil, err
	}

	err = injector.Populate()

	if err != nil {
		return nil, err
	}

	return injector, nil
}

// Resolve resolvs dependency from the injector graph.
func Resolve(injector inject.Graph, name string) (interface{}, error) {
	objects := injector.Objects()

	for i := range objects {
		obj := objects[i]

		if obj.Name == name {
			return obj.Value, nil
		}
	}

	err := fmt.Errorf("dependency %s not found", name)

	return nil, err
}

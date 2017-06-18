package injector

import (
	"fmt"

	"github.com/TsvetanMilanov/todo/server/pkg/handlers"
	"github.com/TsvetanMilanov/todo/server/pkg/services"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/TsvetanMilanov/todo/server/pkg/util"
	"github.com/TsvetanMilanov/todo/server/pkg/validators"
	"github.com/facebookgo/inject"
)

// CreateInjectorGraph registers all required server dependencies and returns their graph.
func CreateInjectorGraph() (*inject.Graph, error) {
	var (
		injector = new(inject.Graph)
		err      error
	)

	var helpers types.IHelpers = &util.Helpers{}
	var dbService types.IDbService = &services.DbService{}
	var serverConfig types.IServerConfig = &services.ServerConfig{}
	var usersService types.IUsersService = &services.UsersService{}
	var modelValidator types.IModelValidator = &validators.ModelValidator{}
	var router types.IRouter = &services.Router{}
	var usersHandler types.IUsersHandler = &handlers.UsersHandler{}
	var authHandler types.IAuthHandler = &handlers.AuthHandler{}
	var authService types.IAuthService = &services.AuthService{}
	var tokensService types.ITokensService = &services.TokensService{}

	server := &services.Server{}

	err = injector.Provide(
		&inject.Object{Value: helpers, Name: "helpers"},
		&inject.Object{Value: server, Name: "server"},
		&inject.Object{Value: dbService, Name: "dbService"},
		&inject.Object{Value: serverConfig, Name: "serverConfig"},
		&inject.Object{Value: usersService, Name: "usersService"},
		&inject.Object{Value: modelValidator, Name: "modelValidator"},
		&inject.Object{Value: router, Name: "router"},
		&inject.Object{Value: usersHandler, Name: "usersHandler"},
		&inject.Object{Value: authHandler, Name: "authHandler"},
		&inject.Object{Value: authService, Name: "authService"},
		&inject.Object{Value: tokensService, Name: "tokensService"},
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

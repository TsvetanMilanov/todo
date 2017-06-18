package injector

import (
	"fmt"

	"github.com/TsvetanMilanov/todo/server/pkg/services"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/TsvetanMilanov/todo/server/pkg/util"
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

	server := &services.Server{}

	err = injector.Provide(
		&inject.Object{Value: helpers, Name: "helpers"},
		&inject.Object{Value: server, Name: "server"},
		&inject.Object{Value: dbService, Name: "dbService"},
		&inject.Object{Value: serverConfig, Name: "serverConfig"},
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

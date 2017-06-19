package injector

import (
	"fmt"

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

	err = injector.Provide(
		&inject.Object{Value: helpers, Name: "helpers"},
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

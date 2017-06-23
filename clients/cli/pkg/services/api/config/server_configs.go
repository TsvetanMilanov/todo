package config

import "github.com/TsvetanMilanov/todo/clients/cli/pkg/types"

// ServerConfigs server configurations.
var ServerConfigs = map[string]types.ServerConfig{
	"local": types.ServerConfig{
		Host:  "localhost:7777",
		Proto: "http",
	},
	"production": types.ServerConfig{
		Host:  "todo-manage.herokuapp.com",
		Proto: "https",
	},
}

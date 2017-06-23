package config

import (
	"errors"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
)

// ServerConfigManager server configuration related methods
type ServerConfigManager struct {
}

// GetServerConfig returns the server configuration for the provided environment.
func (manager *ServerConfigManager) GetServerConfig(env string) (*types.ServerConfig, error) {
	result, ok := ServerConfigs[env]

	if !ok {
		return nil, errors.New("invalid server environment")
	}

	return &result, nil
}

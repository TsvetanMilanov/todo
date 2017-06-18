package services

import "github.com/TsvetanMilanov/todo/server/pkg/types"

// ServerConfig provides methods for configuring the server.
type ServerConfig struct {
	DbService types.IDbService `inject:"dbService"`
}

// Configure configures the server.
func (config *ServerConfig) Configure() error {
	// Configure the DB.
	err := config.DbService.InitializeDatabase()

	return err
}

// Dispose cleans all unwanted resources lefotvers.
func (config *ServerConfig) Dispose() error {
	err := config.DbService.Dispose()

	return err
}

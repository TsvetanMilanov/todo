package types

// IHelpers helper methods.
type IHelpers interface {
	GetEnvVariableOrDefault(env, defaultValue string) string
	EncryptString(value string) (string, error)
}

// IDbService database related methods.
type IDbService interface {
	InitializeDatabase() error
	Dispose() error
}

// IServerConfig server config methods.
type IServerConfig interface {
	Configure() error
	Dispose() error
}

package types

// IHelpers helper methods.
type IHelpers interface {
	GetEnvVariableOrDefault(env, defaultValue string) string
}

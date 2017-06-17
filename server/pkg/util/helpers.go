package util

import (
	"os"
)

// Helpers helper methods
type Helpers struct {
}

// GetEnvVariableOrDefault returns the value of the requested variable or the provided default.
func (helpers *Helpers) GetEnvVariableOrDefault(env, defaultValue string) string {
	result := os.Getenv(env)

	if len(result) == 0 {
		result = defaultValue
	}

	return result
}

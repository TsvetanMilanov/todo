package util

import (
	"os"

	"golang.org/x/crypto/bcrypt"
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

// EncryptString encrypts the provided string.
func (helpers *Helpers) EncryptString(value string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

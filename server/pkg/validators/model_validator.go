package validators

import "fmt"

const (
	usernameMinLength = 4
	usernameMaxLength = 50
	passwordMinLength = 8
	passwordMaxLength = 150
)

// ModelValidator validator for db models.
type ModelValidator struct {
}

// ValidateUser validates the user information.
func (validator *ModelValidator) ValidateUser(username, password string) error {
	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return fmt.Errorf("username should be between %d and %d symbols", usernameMinLength, usernameMaxLength)
	}

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return fmt.Errorf("password should be between %d and %d symbols", passwordMinLength, passwordMaxLength)
	}

	return nil
}

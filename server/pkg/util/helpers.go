package util

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/labstack/echo"

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

// ComparePasswordAndHash compares the provided password with the hashed password.
func (helpers *Helpers) ComparePasswordAndHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GetUserFromContext returns the user which is set in the context.
func (helpers *Helpers) GetUserFromContext(c echo.Context) (*models.User, error) {
	user := c.Get(constants.UserContextVariable)
	if user == nil {
		return nil, echo.ErrUnauthorized
	}

	result, ok := user.(*models.User)
	if !ok {
		return nil, echo.ErrUnauthorized
	}

	return result, nil
}

// GetTokenFromHeader returns the token which is contained in the provided Authorization header.
func (helpers *Helpers) GetTokenFromHeader(headerWithScheme string, authScheme string) (string, error) {
	schemeLength := len(authScheme)

	if len(headerWithScheme) > schemeLength+1 && headerWithScheme[:schemeLength] == authScheme {
		return headerWithScheme[schemeLength+1:], nil
	}

	return "", errors.New("invalid token")
}

// Float64ToInt64 removes the decimal separator from the number.
func (helpers *Helpers) Float64ToInt64(floatNum float64) (int64, error) {
	str := strconv.FormatFloat(floatNum, 'f', 6, 64)
	timestampParts := strings.Split(str, ".")
	return strconv.ParseInt(timestampParts[0]+timestampParts[1], 10, 64)
}

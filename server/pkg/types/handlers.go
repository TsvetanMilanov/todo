package types

import "github.com/labstack/echo"

// IUsersHandler users related methods.
type IUsersHandler interface {
	AddUser(context echo.Context) error
}

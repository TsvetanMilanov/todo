package types

import "github.com/labstack/echo"

// IUsersHandler users related methods.
type IUsersHandler interface {
	AddUser(context echo.Context) error
}

// IAuthHandler auth related methods.
type IAuthHandler interface {
	Login(context echo.Context) error
}

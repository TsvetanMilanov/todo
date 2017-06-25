package types

import "github.com/labstack/echo"

// IAuthMiddleware Auth middlewares
type IAuthMiddleware interface {
	AuthorizeRequest(h echo.HandlerFunc) echo.HandlerFunc
}

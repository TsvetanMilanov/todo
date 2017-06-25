package middlewares

import (
	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/labstack/echo"
)

// Auth middlewares.
type Auth struct {
	Helpers     types.IHelpers     `inject:"helpers"`
	AuthService types.IAuthService `inject:"authService"`
}

// AuthorizeRequest checks if there is Bearer token in the request headers.
func (auth *Auth) AuthorizeRequest(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderValue := c.Request().Header.Get(constants.AuthorizationHeaderName)
		token, err := auth.Helpers.GetTokenFromHeader(authHeaderValue, constants.BearerAuthenticationScheme)
		if err != nil || len(token) == 0 {
			return echo.ErrUnauthorized
		}

		user, err := auth.AuthService.AuthenticateUserWithToken(token)
		if err != nil {
			return echo.ErrUnauthorized
		}

		c.Set(constants.UserContextVariable, user)
		return h(c)
	}
}

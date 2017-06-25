package constants

const (
	// ServerPortValue is the default server port.
	ServerPortValue = "7777"

	// AdminUsername is the username of the site administrator.
	AdminUsername = "administrator"

	// AdminRole the name of the administrator role.
	AdminRole = "administrator"

	// ModeratorRole the name of the moderator role.
	ModeratorRole = "moderator"

	// UserRole the name of the user role.
	UserRole = "user"

	// TokenExpirationTime in secnds.
	TokenExpirationTime = 60 * 60 * 24 * 7

	// BearerTokenType the Bearer token type.
	BearerTokenType = "Bearer"

	// DefaultJwtSecret the default value of the jwt encryption.
	DefaultJwtSecret = "somesecret"

	// UserContextVariable is the name of the context variable which contains the user info.
	UserContextVariable = "User"

	// AuthorizationHeaderName is the name of Authorization header.
	AuthorizationHeaderName = "Authorization"

	// BearerAuthenticationScheme is the name of Bearer authentication scheme.
	BearerAuthenticationScheme = "Bearer"
)

package types

import (
	"net/http"
	"time"

	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

// IHelpers helper methods.
type IHelpers interface {
	GetEnvVariableOrDefault(env, defaultValue string) string
	EncryptString(value string) (string, error)
	ComparePasswordAndHash(password, hashedPassword string) error
	GetUserFromContext(c echo.Context) (*models.User, error)
	GetTokenFromHeader(headerWithScheme string, authScheme string) (string, error)
	Float64ToInt64(floatNum float64) (int64, error)
}

// IDbService database related methods.
type IDbService interface {
	InitializeDatabase() error
	GetCollection(collection string) *mgo.Collection
	Dispose() error
}

// IRouter api routes related methods
type IRouter interface {
	CreateRouter() http.Handler
}

// IServerConfig server config methods.
type IServerConfig interface {
	Configure() error
	Dispose() error
}

// IModelValidator db models validation methods.
type IModelValidator interface {
	ValidateUser(username, password string) error
	ValidateLoginData(username, password string) error
	ValidateNewTodoData(content string, userID string) error
}

// IUsersService describes methods for user related operations.
type IUsersService interface {
	AddUser(username, password string) (*models.User, error)
	GetUser(username string) (*models.User, error)
}

// IAuthService auth related methods
type IAuthService interface {
	Login(username, tokenType string) (*models.Token, error)
	AuthenticateUserWithPassword(username, password string) (*models.User, error)
	AuthenticateUserWithToken(token string) (*models.User, error)
}

// ITokensService token related methods.
type ITokensService interface {
	CreateToken(username, tokenType string) (*models.Token, error)
	ParseToken(token string) (*models.SimpleTokenInfo, error)
	IsTokenExpired(token string) bool
	RemovExpiredTokens(username string) error
}

// ITodosService TODOs related methods.
type ITodosService interface {
	AddTodo(content, category, userID string, deadline *time.Time, priority int) (*models.Todo, error)
}

// IDateService date/time realted methods.
type IDateService interface {
	Now() time.Time
}

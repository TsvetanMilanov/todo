package services

import (
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Router api routes related methods.
type Router struct {
	UsersHandler   types.IUsersHandler   `inject:"usersHandler"`
	AuthHandler    types.IAuthHandler    `inject:"authHandler"`
	TodosHandler   types.ITodosHandler   `inject:"todosHandler"`
	AuthMiddleware types.IAuthMiddleware `inject:"authMiddleware"`
}

// CreateRouter creates the api router.
func (router *Router) CreateRouter() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	api := e.Group("/api")

	router.createUsersGroup(api)
	router.createAuthGroup(api)
	router.createTodosGroup(api)

	api.GET("", func(c echo.Context) error {
		c.JSON(http.StatusOK, "Works")
		return nil
	})

	return e
}

func (router *Router) createUsersGroup(apiGroup *echo.Group) *echo.Group {
	usersGroup := apiGroup.Group("/users")

	usersGroup.POST("", router.UsersHandler.AddUser)

	return usersGroup
}

func (router *Router) createAuthGroup(apiGroup *echo.Group) *echo.Group {
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/login", router.AuthHandler.Login)

	return authGroup
}

func (router *Router) createTodosGroup(apiGroup *echo.Group) *echo.Group {
	todosGroup := apiGroup.Group("/todos", router.AuthMiddleware.AuthorizeRequest)

	todosGroup.POST("", router.TodosHandler.AddTodo)

	return todosGroup
}

package services

import (
	"fmt"
	"net"
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server the API server.
type Server struct {
	Helpers types.IHelpers `inject:"helpers"`
}

// Run starts the server.
func (server *Server) Run() error {
	serverPort := server.Helpers.GetEnvVariableOrDefault(constants.ServerPortEnvVar, constants.ServerPortValue)

	connection, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))

	if err != nil {
		return err
	}

	router := server.createRouter()

	glog.Infof("Server running on port %s", serverPort)
	err = http.Serve(connection, router)

	return err
}

func (server *Server) createRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	api := e.Group("/api")

	api.GET("", func(c echo.Context) error {

		c.JSON(http.StatusOK, "Works")
		return nil
	})

	return e
}

package services

import (
	"fmt"
	"net"
	"net/http"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	"github.com/golang/glog"
)

// Server the API server.
type Server struct {
	Helpers      types.IHelpers      `inject:"helpers"`
	ServerConfig types.IServerConfig `inject:"serverConfig"`
	Router       types.IRouter       `inject:"router"`
}

// Run starts the server.
func (server *Server) Run() error {
	serverPort := server.Helpers.GetEnvVariableOrDefault(constants.ServerPortEnvVar, constants.ServerPortValue)

	connection, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))

	if err != nil {
		return err
	}

	router := server.Router.CreateRouter()

	glog.Infof("Server running on port %s", serverPort)
	err = http.Serve(connection, router)

	return err
}

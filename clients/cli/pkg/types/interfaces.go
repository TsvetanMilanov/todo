package types

import "github.com/spf13/cobra"

// IHelpers helper methods.
type IHelpers interface {
	CheckFlags(cmd *cobra.Command, args []string)
	MarkFlagRequired(cmd *cobra.Command, flag string)
	GetEnv() string
}

// IAuthService authentication related methods.
type IAuthService interface {
	Login(username, password string) (*LoginResponse, error)
}

// IServerConfigManager server configuration related methods.
type IServerConfigManager interface {
	GetServerConfig(env string) (*ServerConfig, error)
}

// IAuthAPIService authentication api related methods.
type IAuthAPIService interface {
	Login(username, password string) (*LoginResponse, error)
}

// IServerClient server api request methods.
type IServerClient interface {
	Post(urlPath string, body interface{}, headers map[string]string, result interface{}) error
}

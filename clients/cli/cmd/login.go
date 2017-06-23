package cmd

import (
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/injector"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var (
	loginCmd = &cobra.Command{
		Use:    "login",
		Short:  "Used to receive access token.",
		Long:   "You will be granted an access token which will be use for operations which require authentication.",
		PreRun: preRun,
		Run:    run,
	}
	username string
	password string
)

func init() {
	rootCmd.AddCommand(loginCmd)
	flagsSet := loginCmd.Flags()

	flagsSet.StringVar(&username, "username", "", "Username")
	flagsSet.StringVar(&password, "password", "", "Password")

	helpers.MarkFlagRequired(loginCmd, "username")
	helpers.MarkFlagRequired(loginCmd, "password")
}

func run(cmd *cobra.Command, args []string) {
	authServiceObj, _ := injector.Resolve(*graph, "authService")
	authService := authServiceObj.(types.IAuthService)

	result, err := authService.Login(username, password)
	if err != nil {
		glog.Exit(err)
	}

	//TODO: Log user data.
	glog.Info(result)
}

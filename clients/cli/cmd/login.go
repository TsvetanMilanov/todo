package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var (
	loginCmd = &cobra.Command{
		Use:    "login",
		Short:  "Used to receive access token.",
		Long:   "You will be granted an access token which will be use for operations which require authentication.",
		PreRun: helpers.CheckFlags,
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
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("login called")
}

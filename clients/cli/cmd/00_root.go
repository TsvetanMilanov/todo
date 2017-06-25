package cmd

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/injector"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
	"github.com/facebookgo/inject"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "TODO CLI client.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Usage()
			if err != nil {
				glog.Error(err)
			}
		},
	}
	graph   *inject.Graph
	helpers types.IHelpers
	logger  types.ILogger
	err     error
)

// Execute ...
func Execute() {
	err := goflag.Set("logtostderr", "true")
	if err != nil {
		glog.Error(err)
	}
	err = goflag.CommandLine.Parse([]string{})
	if err != nil {
		glog.Error(err)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// IMPORTANT dont't change the file name because the init of completion.go
// will be executed before this one. This happens because in the Makefile
// all files are passed in alphabetical order.
func init() {
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	graph, err = injector.CreateInjectorGraph()
	if err != nil {
		glog.Exit(err)
	}

	helpersObject, err := injector.Resolve(*graph, "helpers")
	if err != nil {
		glog.Exit(err)
	}

	helpers = helpersObject.(types.IHelpers)

	loggerObj, err := injector.Resolve(*graph, "logger")
	if err != nil {
		glog.Exit(err)
	}

	logger = loggerObj.(types.ILogger)
}

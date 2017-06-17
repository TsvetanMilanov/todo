package cmd

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "TODO CLI client.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		if err != nil {
			glog.Error(err)
		}
	},
}

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

func init() {
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
}

package util

import (
	"os"
	"strings"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/constants"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Helpers helper methods.
type Helpers struct {
}

// CheckFlags ...
func (helpers *Helpers) CheckFlags(cmd *cobra.Command, args []string) {
	var requiredFlags []string
	var missingFlags []string
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		required, ok := f.Annotations["cobra_annotation_bash_completion_one_required_flag"]
		if ok && required[0] == "true" {
			requiredFlags = append(requiredFlags, f.Name)
		}
	})
	visitedFlags := make(map[string]bool)
	cmd.Flags().Visit(func(f *pflag.Flag) {
		visitedFlags[f.Name] = true
	})

	for _, f := range requiredFlags {
		if _, ok := visitedFlags[f]; !ok {
			missingFlags = append(missingFlags, f)
		}
	}
	if len(missingFlags) == 0 {
		return
	}
	err := cmd.Usage()
	glog.Errorf("Missing flags: %s", strings.Join(missingFlags, ","))
	if err != nil {
		glog.Error(err)
	}
	os.Exit(-1)
}

// MarkFlagRequired ...
func (helpers *Helpers) MarkFlagRequired(cmd *cobra.Command, flag string) {
	var lookupFlag *pflag.Flag
	lookupFlag = cmd.Flags().Lookup(flag)
	if lookupFlag != nil {
		err := cmd.MarkFlagRequired(flag)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}
	err := cmd.MarkPersistentFlagRequired(flag)
	if err != nil {
		glog.Fatal(err)
	}
}

// GetEnv returns the current environment.
func (helpers *Helpers) GetEnv() string {
	env := os.Getenv(constants.EnvVariableName)

	if len(env) == 0 {
		env = constants.LocalEnv
	}

	return env
}

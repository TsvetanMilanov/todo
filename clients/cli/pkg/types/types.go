package types

import "github.com/spf13/cobra"

// IHelpers helper methods.
type IHelpers interface {
	CheckFlags(cmd *cobra.Command, args []string)
	MarkFlagRequired(cmd *cobra.Command, flag string)
}

package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(addcmd)
}

var addcmd = &cobra.Command{
	Use:   "env",
	Short: "Add and remove AWS credentials",
	Long:  "Literally just Add and remove AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

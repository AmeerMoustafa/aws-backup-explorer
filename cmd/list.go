package cmd

import (
	"aws-backup-explorer/app"

	"github.com/spf13/cobra"
)

var runcmd = &cobra.Command{
	Use:   "list",
	Short: "list a DB lookup",
	Long:  "Literally just run a DB lookup",
	Run: func(cmd *cobra.Command, args []string) {
		app.GetBackups()

	},
}

func init() {
	root.AddCommand(runcmd)
}

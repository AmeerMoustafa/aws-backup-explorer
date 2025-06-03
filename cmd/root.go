package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "aws-backup-viewer",
	Version: "1.0",
	Short:   "track Rentworks DB backups",
	Long:    "A simple CLI tool to track today and yesterday's Rentworks database backups stored on AWS",
}

func Execute() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

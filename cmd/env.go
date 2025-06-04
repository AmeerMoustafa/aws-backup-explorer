package cmd

import (
	"aws-backup-explorer/utils"

	"github.com/spf13/cobra"
)

var envcmd = &cobra.Command{
	Use:   "env",
	Short: "Add and remove AWS credentials",
	Long:  "Literally just Add and remove AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {
		accesskey, _ := cmd.Flags().GetString("accesskey")
		secretkey, _ := cmd.Flags().GetString("secretkey")
		utils.SetEnvironment(accesskey, secretkey)

	},
}

func init() {
	root.AddCommand(envcmd)
	envcmd.LocalFlags().String("accesskey", "", "Set your AWS access key")
	envcmd.LocalFlags().String("secretkey", "", "Set your AWS secret key")

}

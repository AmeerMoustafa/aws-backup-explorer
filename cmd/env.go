package cmd

import (
	"aws-backup-explorer/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var envcmd = &cobra.Command{
	Use:   "env",
	Short: "Add and remove AWS credentials",
	Long:  "Literally just Add and remove AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {
		accesskey, err := cmd.Flags().GetString("accesskey")
		if err != nil {
			fmt.Println("[+] The access key flag is REQUIRED")
		}
		secretkey, _ := cmd.Flags().GetString("secretkey")
		utils.SetEnvironment(accesskey, secretkey)

	},
}

var envcheckcmd = &cobra.Command{
	Use:   "check",
	Short: "Check to see whether environment variables exist in the env file",
	Long:  "Literally just Add and remove AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mic check")

	},
}

func init() {
	root.AddCommand(envcmd)
	envcmd.AddCommand(envcheckcmd)
	envcmd.LocalFlags().String("accesskey", "", "Set your AWS access key")
	envcmd.LocalFlags().String("secretkey", "", "Set your AWS secret key")
	envcmd.MarkFlagRequired("accesskey")

}

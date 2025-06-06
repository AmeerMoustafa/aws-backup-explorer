package cmd

import (
	"aws-backup-explorer/utils"
	"fmt"
	"os"

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

var envcheckcmd = &cobra.Command{
	Use:   "check",
	Short: "Check to see whether environment variables exist in the env file",
	Long:  "Literally just Add and remove AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {
		accesskey := os.Getenv("AWS_ACCESS_KEY_ID")
		secretkey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		if accesskey == "" || secretkey == "" {
			fmt.Println("Error: AWS credentials are missing, please add them using the env command")
			return
		} else {

		}

	},
}

var envdelcmd = &cobra.Command{
	Use:   "delete",
	Short: "Clear out your AWS credentials",
	Long:  "Literally just clear out your AWS credentials",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ClearEnvironment()

	},
}

func init() {
	root.AddCommand(envcmd)
	envcmd.AddCommand(envcheckcmd)
	envcmd.AddCommand(envdelcmd)
	envcmd.Flags().String("accesskey", "", "Set your AWS access key")
	envcmd.Flags().String("secretkey", "", "Set your AWS secret key")
	envcmd.MarkFlagRequired("accesskey")
	envcmd.MarkFlagRequired("secretkey")

}

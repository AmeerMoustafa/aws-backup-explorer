package main

import (
	"aws-backup-explorer/cmd"
	"aws-backup-explorer/utils"
	"embed"
)

//go:embed .env
var EnvFile embed.FS

func main() {

	utils.EnvironmentVariables.LoadEnvVariables(EnvFile)

	utils.LoadConfig()

	cmd.Execute()

	// layout := "20060102"

	// t, err := time.Parse(layout, "20250524")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t.Year())

}

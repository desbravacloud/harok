/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package harok

import (
	"fmt"

	"github.com/Spotlitebr/harok/cmd/internal/database"
	"github.com/Spotlitebr/harok/cmd/internal/svn"
	"github.com/spf13/cobra"
)

// createAppCmd represents the createApp command
var createAppCmd = &cobra.Command{
	Use:   "createApp",
	Short: "Create an an application",
	Long: `Create an entry in app table in database.
	This also creates a GitHub repository and a job in Jenkins`,
	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
		}

		hostname, err := cmd.Flags().GetString("hostname")
		if err != nil {
			fmt.Println(err)
		}

		language, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Println(err)
		}

		codeRepo, err := cmd.Flags().GetString("codeRepo")
		if err != nil {
			fmt.Println(err)
		}

		imageRepo, err := cmd.Flags().GetString("imageRepo")
		if err != nil {
			fmt.Println(err)
		}

		var app = new(database.App)
		app.Name = name
		app.Hostname = hostname
		app.Language = language
		app.CodeRepo = codeRepo
		app.ImageRepo = imageRepo

		err = svn.CreateRepo(name, true)
		if err != nil {
			panic(err)
		}

		err = database.InsertIntoAppTable(*app)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("App has been created successfully!!!")
		}
	},
}

func init() {
	rootCmd.AddCommand(createAppCmd)

	createAppCmd.Flags().StringP("name", "n", "", "Application name")
	createAppCmd.Flags().StringP("hostname", "t", "", "Application hostname (used in Virtual Service)")
	createAppCmd.Flags().StringP("language", "l", "", "Application language")
	createAppCmd.Flags().StringP("codeRepo", "r", "", "Application repository")
	createAppCmd.Flags().StringP("imageRepo", "i", "", "Application image repo")

	createAppCmd.MarkFlagRequired("name")
	createAppCmd.MarkFlagRequired("hostname")
	createAppCmd.MarkFlagRequired("language")
	createAppCmd.MarkFlagRequired("codeRepo")
	createAppCmd.MarkFlagRequired("imageRepo")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/services"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"start"},
	Short:   "Create project with specified name and tag from the repo collection",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		tag := cmd.Flag("tag").Value.String()
		name := cmd.Flag("name").Value.String()
		gitInit := cmd.Flag("gitInit").Value.String()
		codeOpen := cmd.Flag("code").Value.String()
		explorerOpen := cmd.Flag("open").Value.String()
		if tag == "" {
			fmt.Println("No tag specified. Creating an empty project")
		}

		if name == "" {
			fmt.Println("No name specified.")
			if tag != "" {
				fmt.Println("Tag name will be used to create the project")
			} else {
				fmt.Println("Error: project name is required!")
				cmd.Help()
				return
			}
		}

		err := services.CreateProject(tag, name, gitInit == "true", codeOpen == "true", explorerOpen == "true")

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().BoolP("help", "h", false, "Show help")
	createCmd.Flags().StringP("tag", "t", "", "Tag of the template to be used")
	createCmd.Flags().StringP("name", "n", "", "Name of the project to be created")
	createCmd.Flags().BoolP("gitInit", "g", false, "Initialize empty git repository")
	createCmd.Flags().BoolP("code", "c", false, "Open project in VS Code")
	createCmd.Flags().BoolP("open", "o", false, "Open project in file explorer")
}

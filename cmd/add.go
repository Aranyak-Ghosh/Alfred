/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/models"
	"alfred/services"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new project(s) to repository collection",
	Long:  `Add new template project(s) to the collection of repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		tag := cmd.Flag("tag").Value.String()
		repo := cmd.Flag("repo").Value.String()
		file := cmd.Flag("file").Value.String()
		branch := cmd.Flag("branch").Value.String()

		overwrite := cmd.Flag("overwrite").Value.String()

		if (tag == "" || repo == "") && file == "" {
			fmt.Println("Error: Missing required flag(s)")
			cmd.Help()
			return
		}
		if tag != "" && repo != "" {
			fmt.Println("Adding project to repository collection...")
			err := services.AddRepoToStore(map[string]models.Repo{tag: {Url: repo, Branch: branch}}, overwrite == "true")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Project added to repository collection")
		}
		if file != "" {
			err := services.AddReposToStoreFromFile(file, overwrite == "true")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Project added to repository collection")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringP("tag", "t", "", "Template Tag used to add project to repository collection")
	addCmd.Flags().StringP("repo", "r", "", "URL of the repository to add")
	addCmd.Flags().StringP("file", "f", "", "File path containing list of repos and tags to be added to collection")
	addCmd.Flags().StringP("branch", "b", "", "Branch to use for the project. If not specified, the default branch will be used")
	addCmd.Flags().BoolP("overwrite", "o", false, "Overwrite existing repository collection")
}

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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the repository for a given tag",
	Long: `Update the repository for a given tag. If the tag does not exist,
the tag will be created if the create flag is set.`,
	Run: func(cmd *cobra.Command, args []string) {
		tag := cmd.Flag("tag").Value.String()
		repo := cmd.Flag("repo").Value.String()
		branch := cmd.Flag("branch").Value.String()
		create := cmd.Flag("create").Value.String()

		if tag == "" || repo == "" {
			fmt.Println("Error: Missing required flag(s)")
			cmd.Help()
			return
		}
		fmt.Printf("Updating %s repository to %s\n", tag, repo)
		err := services.UpdateRepoStore(map[string]models.Repo{tag: {Url: repo, Branch: branch}}, create == "true")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Completed!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().StringP("tag", "t", "", "Template Tag used to update repository collection")
	updateCmd.Flags().StringP("repo", "r", "", "URL of the repository to add")
	updateCmd.Flags().StringP("branch", "b", "default", "Branch to use for the project. If not specified, the default branch will be used")
	updateCmd.Flags().BoolP("create", "c", false, "Create the tag if it does not exist")

}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/services"
	"fmt"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List registered repository",
	Long:    `Get a list of all repositories currently registered in the collection`,
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := services.GetRepoStore()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%-15s %-80s %s\n", "Tag", "Repository", "Branch")
		for tag, repo := range repos {
			fmt.Printf("%-15s %-80s %s\n", tag, repo.Url, repo.Branch)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

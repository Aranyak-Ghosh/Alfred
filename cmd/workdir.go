/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// workdirCmd represents the workdir command
var workdirCmd = &cobra.Command{
	Use:   "workdir",
	Short: "Print execution working directory",
	Long: `Print the directory where the command is called from.

This command can be used to verify that the template will be 
generated in the correct directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := utils.GetWorkingDirectory()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Current working directory: %s\n", dir)
		}

		if cmd.Flag("list").Value.String() == "true" {
			fmt.Println("Listing Directory Contents...")
			files, err := utils.GetDirectoryContents(dir)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				for _, file := range files {
					fmt.Printf("%s\n", file)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(workdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	workdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workdirCmd.Flags().BoolP("list", "l", false, "List directory contents")

}

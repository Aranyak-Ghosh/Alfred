/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"alfred/services"
	"alfred/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alfred",
	Short: "A CLI tool to generate code scaffolding for day to day tasks",
	Long: `Alfred is a CLI tool that can be used to initialize projects 
and potentially configured to set templates which can be used to 
create project scaffolding code and utility tools along with having 
some pre-configured templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Alfred!")
		err := utils.EnsureDependencyInstall()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = services.InitializeRepoStore()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gode-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

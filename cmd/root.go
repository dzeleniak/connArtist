/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "connartist",
	Short: "Create DB conn strings",
	Long: `ConnArtist is a command line application that allows you to quickly generate a Database connection string for your application.
	Version 0.0.1 allows for you to configure your app for local usage with Postgres and cloud usage with MongoDB Atlas.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		
		headerOutput := color.New(color.FgGreen)
		
		headerOutput.Print(`
  _____               ___       __  _     __ 
 / ___/__  ___  ___  / _ | ____/ /_(_)__ / /_
/ /__/ _ \/ _ \/ _ \/ __ |/ __/ __/ (_-</ __/
\___/\___/_//_/_//_/_/ |_/_/  \__/_/___/\__/
		`)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.connartist.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



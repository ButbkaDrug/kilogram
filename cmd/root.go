/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kilogram",
	Short: "Kilogram is a telegram cli tool that created mostly just for fun",
	Long: `Kilogram makes a request to a tdlib server and prints the results
in plain text. It doesn't have a main loop, therefour it's not intendet
to be used like a typical messenging app. But more for is "hacky" way.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
    // Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kilogram.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}



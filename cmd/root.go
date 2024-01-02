package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kilogram",
	Short: "Kilogram is a telegram cli tool that created mostly just for fun",
	Long: `Kilogram makes a request to a tdlib server and prints the results
in plain text. It doesn't have a main loop, therefour it's not intendet
to be used like a typical messenging app. But more for is "hacky" way.`,
    // Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

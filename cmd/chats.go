/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/spf13/cobra"
)

var(
    printAllChats bool
)

// chatsCmd represents the chats command
var chatsCmd = &cobra.Command{
	Use:   "chats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        kg := client.GetChats(printAllChats)
        client.PrintChats(kg.Chats, printAllChats)
	},
}

func init() {
	rootCmd.AddCommand(chatsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    chatsCmd.Flags().BoolVarP(&printAllChats, "all", "a", false, "Show all chats. Chats only with unread messges is default")
}

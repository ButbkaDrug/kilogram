/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

var renderChatsParams *render.RenderChatsParams

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
        kg := client.GetChats()
        renderChatsParams.Order = kg.Positions
        render.RenderChats(kg.Chats, renderChatsParams)
	},
}

func init() {
	getCmd.AddCommand(chatsCmd)

    renderChatsParams = &render.RenderChatsParams{
        PrintAll: false,
        Verbose: false,
    }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    chatsCmd.Flags().BoolVarP(
        &renderChatsParams.PrintAll,
        "all",
        "a",
        false,
        "Show all chats. Chats only with unread messges is default",
    )

    chatsCmd.Flags().BoolVarP(
        &renderChatsParams.Verbose,
        "verbose",
        "v",
        false,
        "Print chat tiltle and unread count along with chat id and last message id",
    )

    chatsCmd.Flags().IntVarP(
        &renderChatsParams.Limit,
        "limit",
        "l",
        0,
        "How many entries you want to print. Defaults to 10",
    )

    chatsCmd.Flags().IntVarP(
        &renderChatsParams.Offset,
        "offset",
        "o",
        0,
        "Specify offset if you want to output chats starting not from first position",
    )
}

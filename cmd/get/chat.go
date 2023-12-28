/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"log"
	"strconv"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        var id int64

        if source >= 0 {
            id = source
        } else if len(args) > 0 {
            i, err := strconv.Atoi(args[0])
            if err != nil {
                log.Fatal(err)
            }

            id = int64(i)
        }

        kc := client.GetChat(id, limit)
        render.PrintChat(kc)
	},
}

func init() {
	getCmd.AddCommand(chatCmd)

    chatCmd.Flags().Int32VarP(
        &limit,
        "limit",
        "l",
        1,
        "Specify number of messages to load",
    )

    chatCmd.Flags().Int64VarP(
        &source,
        "source",
        "s",
        -1,
        "Chat id to be loaded.",
    )
}

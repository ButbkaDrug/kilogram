/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package send

import (
	"log"
	"strings"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/utils"
	"github.com/spf13/cobra"
)

var(
    dest int64
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        var text string

        args = append(args,  utils.ReadStdin()...)


        text = strings.Join(args, " ")

        if len(text) < 1 {
            log.Fatal("Message cannot be empty!")
        }

        client.SendTextMessage(dest, text)
	},
}

func init() {
	sendCmd.AddCommand(messageCmd)

    messageCmd.Flags().Int64VarP(
        &dest,
        "dest",
        "d",
        0,
        "Chat id to send message to.(required)",
    )

    err := messageCmd.MarkFlagRequired("dest")

    if err != nil {
        log.Fatal("Cant mark flag as required: ", err)
    }

}

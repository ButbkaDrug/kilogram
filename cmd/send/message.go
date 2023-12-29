/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package send

import (
	"fmt"
	"os"
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

        stdin, err := utils.ReadStdin()

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }


        args = append(args,  stdin...)
        text = strings.Join(args, "\n")



        msg, err := client.SendTextMessage(dest, text)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        fmt.Println(msg.ChatId, msg.Id)
	},
}

func init() {
	sendCmd.AddCommand(messageCmd)

    messageCmd.Flags().Int64VarP(
        &dest,
        "dest",
        "d",
        0,
        "Chat id to send message to(required). Pass 0 to send message to yourself.",
    )

    if err := messageCmd.MarkFlagRequired("dest"); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}

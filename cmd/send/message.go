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

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Sends a message to specified chat id",
	Long: `You can pipe the text into this command or you can provied text as
    the arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
        var text string

        stdin, err := utils.ReadStdin()

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        text = strings.Join(stdin, "\n")

        if len(stdin) == 0 && len(args) > 0 {
            text = strings.Join(args, "\n")
        }

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

package send

import (
	"fmt"
	"os"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/utils"
	"github.com/spf13/cobra"
)

var SendDocumentParams *client.SendDocumentParams

var documentCmd = &cobra.Command{
	Use:   "document",
	Short: "Sends a file(s) to specified chat id",
    Long: `Use examples:
Sends a file located in a given path to saved messages

    kilogram send document ~/Pictures/image1.jpg`,

	Run: func(cmd *cobra.Command, args []string) {
        stdin, err := utils.ReadStdin()

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if len(stdin) > 0 {
            args = stdin
        }

        SendDocumentParams.Files = append(SendDocumentParams.Files, args...)

        msgs, err := client.SendDocument(SendDocumentParams)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        fmt.Println(msgs.ClientId, msgs.Messages)
	},
}

func init() {
	sendCmd.AddCommand(documentCmd)

    SendDocumentParams= &client.SendDocumentParams{
        ChatId: 0,
        Files: []string{},
        Caption: "",
    }

    documentCmd.Flags().Int64VarP(
        &SendDocumentParams.ChatId,
        "dest",
        "d",
        0,
        "Id of a destination chat. If no destination provided message will be sent to saved messages",
    )

    documentCmd.Flags().StringVarP(
        &SendDocumentParams.Caption,
        "caption",
        "c",
        "",
        "Caption tha describes file being send(Optional)",
    )
}

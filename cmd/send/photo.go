package send

import (
	"fmt"
	"os"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/utils"
	"github.com/spf13/cobra"
)

var SendPhotoParams*client.SendPhotoParams

var photoCmd = &cobra.Command{
	Use:   "photo",
	Short: "Sends a image file(s) to specified chat id",
    Long: `Use examples:
Sends a photo located in a given path to saved messages

    kilogram send photo ~/Pictures/image1.jpg`,

	Run: func(cmd *cobra.Command, args []string) {
        stdin, err := utils.ReadStdin()

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        args = append(args, stdin...)

        fmt.Println("Piped args: ", args)
        SendPhotoParams.Files = append(SendPhotoParams.Files, args...)

        msgs, err := client.SendPhoto(SendPhotoParams)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        fmt.Println(msgs.ClientId, msgs.Messages)
	},
}

func init() {
	sendCmd.AddCommand(photoCmd)

    SendPhotoParams= &client.SendPhotoParams{
        ChatId: 0,
        Files: []string{},
        Caption: "",
        Width: 0,
        Height: 0,
        Spoiler: false,
    }

    photoCmd.Flags().Int64VarP(
        &SendPhotoParams.ChatId,
        "dest",
        "d",
        0,
        "Id of a destination chat. If no destination provided message will be sent to saved messages",
    )

    photoCmd.Flags().Int32VarP(
        &SendPhotoParams.Width,
        "width",
        "w",
        0,
        "Width of the picture",
    )

    photoCmd.Flags().Int32VarP(
        &SendPhotoParams.Height,
        "height",
        "e",
        0,
        "Height of a picture",
    )

    photoCmd.Flags().BoolVar(
        &SendPhotoParams.Spoiler,
        "hide",
        false,
        "Set this flag if photo needs to be hidden under a spoiler",
    )


    photoCmd.Flags().StringVarP(
        &SendPhotoParams.Caption,
        "caption",
        "c",
        "",
        "Caption tha describes file being send(Optional)",
    )
}

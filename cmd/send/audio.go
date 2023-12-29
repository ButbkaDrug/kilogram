/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package send

import (
	"fmt"
	"os"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/spf13/cobra"
)

var SendAudioParams *client.SendAudioParams

// audioCmd represents the audio command
var audioCmd = &cobra.Command{
	Use:   "audio",
	Short: "Sends an audio file to specified chat id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        client.SendAudio(SendAudioParams)
	},
}

func init() {
	sendCmd.AddCommand(audioCmd)

    SendAudioParams = &client.SendAudioParams{
        Id: 0,
        File: "",
        Caption: "",
    }

    audioCmd.Flags().Int64VarP(
        &SendAudioParams.Id,
        "dest",
        "d",
        0,
        "Id of a destination chat. If no destination provided message will be sent to saved messages",
    )


    audioCmd.Flags().StringVarP(
        &SendAudioParams.File,
        "file",
        "f",
        "",
        "path to an audio file to be send(Required)",
    )

    audioCmd.Flags().StringVarP(
        &SendAudioParams.Caption,
        "caption",
        "c",
        "",
        "Caption tha describes file being send(Optional)",
    )

    if err := audioCmd.MarkFlagRequired("file"); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

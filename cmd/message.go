/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/spf13/cobra"
)

var(
    file string
    text string
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

        text = strings.Join(args, " ")

        if file != "" {

            bytes, err := os.ReadFile(file)

            if err != nil {
                log.Fatal(err)
            }

            text = string(bytes)

        }

        if len(text) < 1 {
            log.Fatal("Message cannot be empty!")
        }
        client.SendTextMessage(dest, text)
	},
}

func init() {
	newCmd.AddCommand(messageCmd)

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

    messageCmd.Flags().StringVarP(
        &file,
        "file",
        "f",
        "",
        "Path to text file to send as a message",
    )

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// messageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

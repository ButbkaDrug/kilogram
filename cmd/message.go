/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/spf13/cobra"
)

var(
    file string
    msg string
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

        if msg != "" {
            text = msg
        } else if file != "" {
            text = readFromFile(file)
        } else if m := readFromStdin(); m != "" {
            text = m
        } else if len(args) > 0 {
            text = strings.Join(args, " ")
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

    messageCmd.Flags().StringVarP(
        &msg,
        "message",
        "m",
        "",
        "Message to be sent.",
    )

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// messageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func readFromStdin() string {

        var text string

        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {

            text += scanner.Text()
            text += "\n"

        }

        return text
}

func readFromFile(fp string) string {

        var text string

        bytes, err := os.ReadFile(fp)

        if err != nil {
            log.Fatal(err)
        }

        text = string(bytes)


    return text
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/spf13/cobra"
)

var (
    source int64
    dest int64
    messageIds []int64
)

// forwardCmd represents the forward command
var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        for _, arg := range args {
            id, err := strconv.Atoi(arg)

            if err != nil {
                log.Fatal("Argument parsing error: ", err)
            }

            messageIds = append(messageIds, int64(id))
        }
        client.ForwardMessage(source, dest, messageIds)
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)

    forwardCmd.Flags().Int64VarP(
        &source,
        "source",
        "s",
        0,
        "Original chat id messages belong to.(required)",
    )

    if err := forwardCmd.MarkFlagRequired("source"); err != nil {
        log.Fatal(err)
    }
    forwardCmd.Flags().Int64VarP(
        &dest,
        "dest",
        "d",
        0,
        "Chat id for messages to be forwarded to.(required)",
    )

    if err := forwardCmd.MarkFlagRequired("dest"); err != nil {
        log.Fatal(err)
    }




	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"
	"strings"

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
	Short: "Forward message(s) from one chat to another by their ids",
	Long: `Use this function to forward message(s) to -dest chat id. Pass 0 as
dest to forward message(s) to yourself. If source not specified, first argumen
will be treated as a source id, to make chaning with other commands easier.

If no message ids specified, last message in the chat will be forwarded

For example:

forward -s 123456789 -d 0 111 222 333

Will forward messages with ids 111, 222 and 333 from chat with id 123456789 to Saved Messages`,

	Run: func(cmd *cobra.Command, args []string) {


        if pipe := readFromStdin(); pipe != "" {
            pargs := strings.Split(pipe, " ")
            args = append(args, pargs...)
        }


        ids, err := argsToIds(args)

        if err != nil {
            log.Fatal("Error converting args: ", err)
        }

        if source == -1 {
            source = ids[0]
            ids = ids[1:]
        }

        client.ForwardMessage(&client.ForwardMessageParams{
            Source: source,
            Dest: dest,
            Messages: ids,
            Limit: limit,
        })
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)

    forwardCmd.Flags().Int64VarP(
        &source,
        "source",
        "s",
        -1,
        "Original chat id messages belong to.(required)",
    )


    forwardCmd.Flags().Int64VarP(
        &dest,
        "dest",
        "d",
        -1,
        "Chat id for messages to be forwarded to.(required)",
    )

    if err := forwardCmd.MarkFlagRequired("dest"); err != nil {
        log.Fatal(err)
    }

    forwardCmd.Flags().Int32VarP(
        &limit,
        "limit",
        "l",
        1,
        "If no message IDs provided. How many messages to forward, starting from the last message fround in the chat",
    )


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func argsToIds(args []string) ([]int64, error) {
    var ids []int64
        for _, arg := range args {
            arg = strings.Trim(arg, "\n\t")
            if arg == "" { continue }
            id, err := strconv.Atoi(arg)

            if err != nil {
                return ids, err
            }

            ids = append(ids, int64(id))
        }
    return ids, nil
}

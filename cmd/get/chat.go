package get

import (
	"strconv"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Loads history of a specified chat",
	Long: `If you don't provied chat id, saved messages will be loaded
    You can also provide -l flag to specify how many messages starting from the
    last one you want to load.`,
	Run: func(cmd *cobra.Command, args []string) {
        var id int64

        if source >= 0 {
            id = source
        } else if len(args) > 0 {
            i, err := strconv.Atoi(args[0])
            if err != nil {
                panic(err)
            }

            id = int64(i)
        }

        kc := client.GetChat(id, limit)
        render.PrintChat(kc)

	},
}

func init() {
	getCmd.AddCommand(chatCmd)

    chatCmd.Flags().Int32VarP(
        &limit,
        "limit",
        "l",
        1,
        "Specify number of messages to load",
    )

    chatCmd.Flags().Int64VarP(
        &source,
        "source",
        "s",
        -1,
        "Chat id to be loaded.",
    )
}

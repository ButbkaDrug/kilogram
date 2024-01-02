package get

import (
	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

var renderChatsParams *render.RenderChatsParams

var chatsCmd = &cobra.Command{
	Use:   "chats",
	Short: "Loads chats from your main chat list",
	Long: `By default load chats will display chats with unread messages only.
Pass --all flag to load all the chats. You can also specify how many chats you
want to load by passing -l flag.
EXAMPLE:
kilogram load chats -a -v -l 20 - will load first 20 chats from a main chat list`,
	Run: func(cmd *cobra.Command, args []string) {
        kg := client.GetChats()
        renderChatsParams.Order = kg.Positions
        render.RenderChats(kg.Chats, renderChatsParams)
	},
}

func init() {
	getCmd.AddCommand(chatsCmd)

    renderChatsParams = &render.RenderChatsParams{
        PrintAll: false,
        Verbose: false,
    }

    chatsCmd.Flags().BoolVarP(
        &renderChatsParams.PrintAll,
        "all",
        "a",
        false,
        "Show all chats. Chats only with unread messges is default",
    )

    chatsCmd.Flags().BoolVarP(
        &renderChatsParams.Verbose,
        "verbose",
        "v",
        false,
        "Print chat tiltle and unread count along with chat id and last message id",
    )

    chatsCmd.Flags().IntVarP(
        &renderChatsParams.Limit,
        "limit",
        "l",
        0,
        "How many entries you want to print. Defaults to 10",
    )

    chatsCmd.Flags().IntVarP(
        &renderChatsParams.Offset,
        "offset",
        "o",
        0,
        "Specify offset if you want to output chats starting not from first position",
    )
}

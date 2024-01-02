package get

import (
    "github.com/butbkadrug/kilogram/cmd"
	"github.com/spf13/cobra"
)

var(
    source int64
    limit int32
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get command will help you to retrive information from telegram",
	Long: `Try "get chat" or "get users"`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}

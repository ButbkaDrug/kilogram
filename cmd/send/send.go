package send

import (
    "github.com/butbkadrug/kilogram/cmd"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Will help you send files and messages",
    Long: `For now you can send audio, pictures and text messages`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	cmd.RootCmd.AddCommand(sendCmd)
}

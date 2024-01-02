package get

import (
	"log"
	"strconv"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)
var renderUserParams *render.RenderUserParams

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Gets informatin about specified user id",
    Args: cobra.MinimumNArgs(1),
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.ParseInt(args[0], 10, 64)

        if err != nil {
            log.Fatal("Invalid user id", err)
        }


        user := client.GetUser(id)
        render.RenderUser(user, renderUserParams)
	},
}

func init() {
	getCmd.AddCommand(userCmd)


    renderUserParams = &render.RenderUserParams{
        Verbose: true,
        Status: false,
    }

    userCmd.Flags().BoolVarP(
        &renderUserParams.Verbose,
        "verbose",
        "v",
        true,
        "Prittyfied output. Containing all avalible info",
    )

    userCmd.Flags().BoolVarP(
        &renderUserParams.Status,
        "status",
        "t",
        false,
        "Output user status.",
    )
}

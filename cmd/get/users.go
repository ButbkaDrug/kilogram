package get

import (

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

var renderUsersParams *render.RenderUsersParams

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Will load userf from your contact list",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        users := client.GetUsers()
        render.RenderUsers(users, renderUsersParams)
	},
}

func init() {
	getCmd.AddCommand(usersCmd)

    renderUsersParams = &render.RenderUsersParams{
        Verbose: true,
    }

    usersCmd.Flags().BoolVarP(
        &renderUsersParams.All,
        "all",
        "a",
        false,
        "List all contacts(defauil only to online)",
    )
}

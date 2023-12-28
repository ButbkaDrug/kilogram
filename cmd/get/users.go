/* Copyright © 2023 NAME HERE <EMAIL ADDRESS> */
package get

import (

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)

var renderUsersParams *render.RenderUsersParams

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

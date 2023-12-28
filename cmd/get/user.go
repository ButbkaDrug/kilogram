/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"log"
	"strconv"

	"github.com/butbkadrug/kilogram/internal/client"
	"github.com/butbkadrug/kilogram/internal/render"
	"github.com/spf13/cobra"
)
var renderUserParams *render.RenderUserParams

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
    Args: cobra.MinimumNArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

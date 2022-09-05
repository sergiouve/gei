package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable an installed extensions",
	Run: func(cmd *cobra.Command, args []string) {
		enable(args)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}

func enable(args []string) {
	if (len(args) < 1) {
		fmt.Println("An extension UUID is required")
		os.Exit(0)
	}

	error := installer.EnableExtension(args[0])

	if error != nil {
		fmt.Println("Error enabling extension: ", error.Error())
	} else {
		fmt.Println("Extension enabled")
	}
}

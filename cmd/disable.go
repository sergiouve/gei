package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an installed extensions",
	Run: func(cmd *cobra.Command, args []string) {
		disable(args)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}

func disable(args []string) {
	if (len(args) < 1) {
		fmt.Println("An extension UUID is required")
		os.Exit(0)
	}

	error := installer.DisableExtension(args[0])

	if error != nil {
		fmt.Println("Error enabling extension: ", error.Error())
	} else {
		fmt.Println("Extension disabled")
	}
}

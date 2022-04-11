package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove installed extensions",
	Run: func(cmd *cobra.Command, args []string) {
		remove(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remove(args []string) {
	if (len(args) < 1) {
		fmt.Println("An id or list of ids are required")
		os.Exit(0)
	}

	installer.RemoveExtension()
}

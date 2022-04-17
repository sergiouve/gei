package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh local database",
	Run: func(cmd *cobra.Command, args []string) {
		refresh(args)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}

func refresh(args []string) {
	installer.RefreshLocalDatabase()
}

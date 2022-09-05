package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed extensions",
	Run: func(cmd *cobra.Command, args []string) {
		list(args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(args []string) {
	listResult := installer.ListInstalledExtensions()

	for _, result := range listResult {
		fmt.Println(fmt.Sprintf("uuid: %s, url: %s", result.Uuid, result.DownloadUrl))
	}
}

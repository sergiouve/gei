package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search available extensions in https://extensions.gnome.org",
	Run: func(cmd *cobra.Command, args []string) {
		search(args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(args []string) {
	if (len(args) < 1) {
		fmt.Println("A search term is required")
		os.Exit(0)
	}

	searchResults := installer.FetchSearch(args[0])

	for _, result := range searchResults.Extensions {
		fmt.Println(fmt.Sprintf("id: %d, name: %s, url: https://extensions.gnome.org%s", result.Pk, result.Name, result.Link))
	}
}

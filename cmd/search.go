package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"gitlab.com/yugarinn/gei/installer"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search available extensions in https://extensions.gnome.org",
	Long: ` A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		search(args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(args []string) {
	searchResults := installer.FetchSearch(args[0])

	for _, result := range searchResults {
		fmt.Println(result.Pk)
		fmt.Println(result.Name)
		fmt.Println(result.Description)
		fmt.Println("---")
	}
}

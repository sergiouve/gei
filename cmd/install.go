package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/yugarinn/gei/installer"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install available extensions from https://extensions.gnome.org",
	Run: func(cmd *cobra.Command, args []string) {
		install(args)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(args []string) {
	if (len(args) < 1) {
		fmt.Println("An extension ID is required")
		os.Exit(0)
	}

	installer.InstallExtension(args[0])

	fmt.Println("Extension successfully installed")
}

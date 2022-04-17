package installer

import (
	"fmt"
	"os"
	"gitlab.com/yugarinn/gei/installer/idos"
	"github.com/godbus/dbus/v5"
)

func ListInstalledExtensions() []idos.ExtensionMetadata {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var rawInstalledExtensions map[string]map[string]dbus.Variant
	err = conn.Object("org.gnome.Shell", "/org/gnome/Shell").Call("org.gnome.Shell.Extensions.ListExtensions", 0).Store(&rawInstalledExtensions)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get list of owned names:", err)
		os.Exit(1)
	}

	var parsedInstalledExtensions []idos.ExtensionMetadata

	for uuid := range rawInstalledExtensions {
		var extensionMetadata idos.ExtensionMetadata

		extensionMetadata.Uuid = rawInstalledExtensions[uuid]["uuid"].String()
		extensionMetadata.DownloadUrl = rawInstalledExtensions[uuid]["url"].String()

		parsedInstalledExtensions = append(parsedInstalledExtensions, extensionMetadata)
	}

	return parsedInstalledExtensions
}

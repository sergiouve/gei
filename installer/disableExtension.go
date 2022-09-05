package installer

import (
	"github.com/godbus/dbus/v5"
)

func DisableExtension(extensionUuid string) error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Object("org.gnome.Shell", "/org/gnome/Shell").Call("org.gnome.Shell.Extensions.DisableExtension", 0, extensionUuid)

	return nil
}

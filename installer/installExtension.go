package installer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gitlab.com/yugarinn/gei/installer/client"
	"gitlab.com/yugarinn/gei/installer/idos"
	"github.com/godbus/dbus/v5"
)

func InstallExtension(extensionId string) error {
	extensionMetadata, err := getExtensionMetadata(extensionId)

	downloadExtension(extensionMetadata)
	UnzipExtension(extensionMetadata.Uuid)
	enableExtension(extensionMetadata.Uuid)
	deleteZip(extensionMetadata.Uuid)

	return err
}

func getExtensionMetadata(extensionId string) (idos.ExtensionMetadata, error) {
	systemShellVersion, err := getSystemShellMajorVersion()
	extensionMetadataResponse, err := client.FetchExtensionMetadata(extensionId, systemShellVersion)

	var extensionMetadata idos.ExtensionMetadata
	json.Unmarshal(extensionMetadataResponse, &extensionMetadata)

	return extensionMetadata, err
}

func getSystemShellMajorVersion() (string, error) {
	rawShellCommandOutput, err := exec.Command("gnome-shell", "--version").Output()

	splittedCommandOutput := strings.Split(string(rawShellCommandOutput), " ")
	gnomeShellVersion := splittedCommandOutput[len(splittedCommandOutput) - 1]
	gnomeShellMajorVersion := strings.Split(gnomeShellVersion, ".")[0]

	return gnomeShellMajorVersion, err
}

func downloadExtension(extensionMetadata idos.ExtensionMetadata) {
	client.DownloadExtension(extensionMetadata)
}

func enableExtension(extensionUuid string) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	conn.Object("org.gnome.Shell", "/org/gnome/Shell").Call("org.gnome.Shell.Extensions.EnableExtension", 0, extensionUuid)
}

func deleteZip(uuid string) {
	homeDir, _ := os.UserHomeDir()
	fileName := fmt.Sprintf("%s.zip", uuid)

	os.Remove(filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(fileName)))
}

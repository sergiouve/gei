package installer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gitlab.com/yugarinn/gei/client"
	"gitlab.com/yugarinn/gei/idos"
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
	err2 := exec.Command("gnome-extensions", "enable", extensionUuid).Run()
	err3 := exec.Command("dbus-send", "--session", "--type=method_call", "--dest=org.gnome.Shell /org/gnome/Shell org.gnome.Shell.Eval string:\"global.reexec_self();\"").Run()

	if err2 != nil {
		fmt.Println(err2)
	}

	if err3 != nil {
		fmt.Println(err3)
	}
}


func deleteZip(uuid string) {
	homeDir, _ := os.UserHomeDir()
	fileName := fmt.Sprintf("%s.zip", uuid)

	os.Remove(filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(fileName)))
}

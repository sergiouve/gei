package installer

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"path/filepath"
	"os"
	"os/exec"
	"io"
	"strings"

	"gitlab.com/yugarinn/gei/client"
	"gitlab.com/yugarinn/gei/idos"
)

func InstallExtension(extensionId string) {
	extensionMetadata := getExtensionMetadata(extensionId)

	downloadExtension(extensionMetadata)
	unzipExtension(extensionMetadata.Uuid)
	deleteZip(extensionMetadata.Uuid)
}

func getExtensionMetadata(extensionId string) idos.ExtensionMetadata {
	systemShellVersion := getSystemShellMajorVersion()
	extensionMetadataResponse := client.FetchExtensionMetadata(extensionId, systemShellVersion)

	var extensionMetadata idos.ExtensionMetadata
	json.Unmarshal(extensionMetadataResponse, &extensionMetadata)

	return extensionMetadata
}

func getSystemShellMajorVersion() string {
	rawShellCommandOutput, err := exec.Command("gnome-shell", "--version").Output()

	if err != nil {
		fmt.Println(err)
	}

	splittedCommandOutput := strings.Split(string(rawShellCommandOutput), " ")
	gnomeShellVersion := splittedCommandOutput[len(splittedCommandOutput) - 1]
	gnomeShellMajorVersion := strings.Split(gnomeShellVersion, ".")[0]

	return gnomeShellMajorVersion
}

func downloadExtension(extensionMetadata idos.ExtensionMetadata) {
	client.DownloadExtension(extensionMetadata)
}

func unzipExtension(uuid string) error {
	homeDir, _ := os.UserHomeDir()
	fileName := fmt.Sprintf("%s.zip", uuid)
	dest := filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(uuid))

    r, err := zip.OpenReader(filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(fileName)))

    if err != nil {
        return err
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()

    os.MkdirAll(dest, 0755)

    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        if !strings.HasPrefix(path, filepath.Clean(dest) + string(os.PathSeparator)) {
            return fmt.Errorf("illegal file path: %s", path)
        }

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            os.MkdirAll(filepath.Dir(path), f.Mode())
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }

    return nil
}

func deleteZip(uuid string) {
	homeDir, _ := os.UserHomeDir()
	fileName := fmt.Sprintf("%s.zip", uuid)

	os.Remove(filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(fileName)))
}

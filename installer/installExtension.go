package installer

import (
	"gitlab.com/yugarinn/gei/client"
	"encoding/json"
	"fmt"
)

type ExtensionMetadata struct {
	DownloadUrl string `json:"download_url"`
}

func InstallExtension(extensionId string) {
	extensionMetadata := getExtensionMetadata(extensionId)

	fmt.Println(extensionMetadata)
}

func getExtensionMetadata(extensionId string) ExtensionMetadata {
	systemShellVersion := getSystemShellVersion()
	extensionMetadataResponse := client.FetchExtensionMetadata(extensionId, systemShellVersion)

	var extensionMetadata ExtensionMetadata
	json.Unmarshal(extensionMetadataResponse, &extensionMetadata)

	return extensionMetadata
}

func getSystemShellVersion() string {
	return "41"
}

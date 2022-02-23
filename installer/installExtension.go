package installer

import (
	"gitlab.com/yugarinn/gei/client"
)

func InstallExtension(extensionId string) {
	client.DownloadExtension(extensionId)
}

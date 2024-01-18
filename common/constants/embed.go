package constants

import "embed"

var (
	EmbedWebappPath string
	EmbedWebappFS   embed.FS
)

func SetEmbed(path string, fs embed.FS) {
	EmbedWebappPath = path
	EmbedWebappFS = fs
}

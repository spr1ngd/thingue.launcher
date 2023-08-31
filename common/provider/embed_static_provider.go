package provider

import "embed"

var (
	WebStaticPath  string
	WebStaticFiles embed.FS
)

func SetWebStatic(path string, fs embed.FS) {
	WebStaticPath = path
	WebStaticFiles = fs
}

package web

import "embed"

//go:embed client
var StaticFiles embed.FS

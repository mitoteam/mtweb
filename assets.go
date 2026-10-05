package mtweb

import (
	"embed"
	"io/fs"
)

// embedded web assets
//
//go:embed assets/vendor/*.css assets/vendor/*.js
//go:embed assets/webfonts/*
var embedFS embed.FS

var MtWebAssetsFS fs.FS

func init() {
	//prepare fs.FS for embedded subdirectory
	MtWebAssetsFS, _ = fs.Sub(embedFS, "assets")
}

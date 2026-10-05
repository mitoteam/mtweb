package mtweb

import (
	"embed"
	"io/fs"

	"github.com/mitoteam/mbr"
)

// embedded web assets
//
//go:embed assets/vendor/*.css assets/vendor/*.js
//go:embed assets/webfonts/*
var embedFS embed.FS

var webAssetsFS fs.FS

type AssetsRouteControllerType struct {
	mbr.ControllerBase
}

var assetsRouteController *AssetsRouteControllerType

// This route should be added to root route controller. Example:
// func (c *RootController) MtWebAssets() mbr.Route { return mtweb.AssetsRoute }
var AssetsRoute mbr.Route

func init() {
	//prepare fs.FS for embedded subdirectory
	webAssetsFS, _ = fs.Sub(embedFS, "assets")

	//prepare route controller to use as child controller
	assetsRouteController = &AssetsRouteControllerType{}

	AssetsRoute = mbr.Route{PathPattern: "/assets/mtweb", ChildController: assetsRouteController}
}

func (c *AssetsRouteControllerType) MtWebAssets() mbr.Route {
	return mbr.Route{PathPattern: "/", StaticFS: webAssetsFS}
}

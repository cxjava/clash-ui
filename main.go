package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed yacd/*
var yacd embed.FS

//go:embed clash/*
var clash embed.FS

var addr = flag.String("l", ":8088", "Listen address")
var yacdPath = flag.String("y", "/y", "https://github.com/haishanh/yacd path")
var clashPath = flag.String("c", "/c", "https://github.com/Dreamacro/clash-dashboard path ")

func main() {
	flag.Parse()
	r := gin.Default()

	// Go's `embed.FS` preserves relative paths, so paths within `fs` contain the relative path `static`.
	// The solution is simple: use the `Sub()` function in `io/fs` to return the filesystem of the desired subdirectory.
	yacdRoot, _ := fs.Sub(yacd, "yacd")
	r.StaticFS(*yacdPath, http.FS(yacdRoot))

	clashRoot, _ := fs.Sub(clash, "clash")
	r.StaticFS(*clashPath, http.FS(clashRoot))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, *yacdPath)
	})
	r.Run(*addr)
}

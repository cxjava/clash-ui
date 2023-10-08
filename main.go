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

func main() {
	r := gin.Default()

	// Go 的 `embed.FS` 保留了相对路径，因此 `fs` 内的路径是包含相对路径 `static` 的。
	// 解决办法也很简单，使用 `io/fs` 里的 `Sub()` 函数返回所需的子目录的文件系统即可。
	yacdRoot, _ := fs.Sub(yacd, "yacd")
	r.StaticFS("/yacd", http.FS(yacdRoot))

	clashRoot, _ := fs.Sub(clash, "clash")
	r.StaticFS("/ui", http.FS(clashRoot))

	r.Run(*addr)
}

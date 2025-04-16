package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve YACD static files
	yacdRoot, _ := fs.Sub(yacd, "yacd")
	yacdHandler := http.StripPrefix(*yacdPath, http.FileServer(http.FS(yacdRoot)))
	r.Handle(*yacdPath, http.RedirectHandler(*yacdPath+"/", http.StatusMovedPermanently))
	r.Handle(*yacdPath+"/*", yacdHandler)

	// Serve Clash static files
	clashRoot, _ := fs.Sub(clash, "clash")
	clashHandler := http.StripPrefix(*clashPath, http.FileServer(http.FS(clashRoot)))
	r.Handle(*clashPath, http.RedirectHandler(*clashPath+"/", http.StatusMovedPermanently))
	r.Handle(*clashPath+"/*", clashHandler)
	// Redirect root to YACD
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *yacdPath, http.StatusMovedPermanently)
	})

	// Start server
	http.ListenAndServe(*addr, r)
}

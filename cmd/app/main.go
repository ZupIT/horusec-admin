package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
	"github.com/tiagoangelozup/horusec-admin/internal/http/router"
	"github.com/tiagoangelozup/horusec-admin/internal/server"
)

var rnd = renderer.New(renderer.Options{ParseGlobPattern: "web/template/*.gohtml"})

func main() {
	r, err := router.New()
	checkErr(err)

	// static files
	dir, _ := os.Getwd()
	fileServer(r, "/icons", http.Dir(filepath.Join(dir, "web/static/icons")))
	fileServer(r, "/images", http.Dir(filepath.Join(dir, "web/static/images")))
	fileServer(r, "/scripts", http.Dir(filepath.Join(dir, "web/static/scripts")))
	fileServer(r, "/styles", http.Dir(filepath.Join(dir, "web/static/styles")))

	// error page
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := rnd.HTML(w, http.StatusNotFound, "not-found", nil)
		checkErr(err)
	})

	srv := server.New(r).Start()

	waitForInterruptSignal()
	err = srv.GracefullyShutdown()
	if err != nil {
		log.Fatal(fmt.Errorf("server forced to shutdown: %w", err))
	}

	log.Println("server exiting")
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		prefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(prefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

var rnd = renderer.New(renderer.Options{ParseGlobPattern: "web/template/*.gohtml"})

const addr = ":3001"

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := rnd.HTML(w, http.StatusOK, "index", nil)
		checkErr(err)
	})

	dir, _ := os.Getwd()
	fileServer(r, "/icons", http.Dir(filepath.Join(dir, "web/static/icons")))
	fileServer(r, "/images", http.Dir(filepath.Join(dir, "web/static/images")))
	fileServer(r, "/scripts", http.Dir(filepath.Join(dir, "web/static/scripts")))
	fileServer(r, "/styles", http.Dir(filepath.Join(dir, "web/static/styles")))

	srv := &http.Server{Addr: addr, Handler: r}
	go func() {
		log.Println("Listening on ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
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
		log.Fatal(err) //respond with error page or message
	}
}

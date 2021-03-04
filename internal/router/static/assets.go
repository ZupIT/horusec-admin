package static

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

const relativePath = "web/static"

type (
	directory struct {
		Pattern string
		http.FileSystem
	}
	Assets []*directory
)

func ListAssets() (Assets, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to find the assets path: %w", err)
	}
	dir := path.Join(wd, relativePath)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var folders []*directory
	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, &directory{Pattern: "/" + f.Name(), FileSystem: http.Dir(filepath.Join(dir, f.Name()))})
		}
	}

	return folders, nil
}

func (f *directory) Serve(r *chi.Mux) {
	if strings.ContainsAny(f.Pattern, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if f.Pattern != "/" && f.Pattern[len(f.Pattern)-1] != '/' {
		r.Get(f.Pattern, http.RedirectHandler(f.Pattern+"/", http.StatusMovedPermanently).ServeHTTP)
		f.Pattern += "/"
	}
	f.Pattern += "*"

	r.Get(f.Pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		prefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(prefix, http.FileServer(f))
		fs.ServeHTTP(w, r)
	})
}

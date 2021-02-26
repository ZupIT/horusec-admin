package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const assetsPath = "/web/static"

type Asset struct {
	Pattern   string
	Directory http.FileSystem
}

func scanAssets() ([]*Asset, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to find the assets path: %w", err)
	}
	dir := path.Join(wd, assetsPath)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var assets []*Asset
	for _, f := range files {
		if f.IsDir() {
			assets = append(assets, &Asset{Pattern: "/" + f.Name(), Directory: http.Dir(filepath.Join(dir, f.Name()))})
		}
	}
	return assets, nil
}

func (a *Asset) serve(r *chi.Mux) {
	if strings.ContainsAny(a.Pattern, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if a.Pattern != "/" && a.Pattern[len(a.Pattern)-1] != '/' {
		r.Get(a.Pattern, http.RedirectHandler(a.Pattern+"/", http.StatusMovedPermanently).ServeHTTP)
		a.Pattern += "/"
	}
	a.Pattern += "*"

	r.Get(a.Pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		prefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(prefix, http.FileServer(a.Directory))
		fs.ServeHTTP(w, r)
	})
}

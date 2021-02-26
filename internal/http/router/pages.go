package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/thedevsaddam/renderer"
)

const pagesPath = "/web/template/pages"

type Page struct {
	render *renderer.Render
	file   string

	Name    string
	Pattern string
}

func newPage(render *renderer.Render, file string, dir string) *Page {
	pattern := strings.TrimPrefix(file, dir)
	pattern = strings.TrimSuffix(pattern, filepath.Ext(file))
	return &Page{
		render: render,
		file:   file,
		Name:   strings.TrimPrefix(pattern, "/"),
		Pattern: func() string {
			if pattern == "/index" {
				return "/"
			}
			return pattern
		}(),
	}
}

func scanPages(render *renderer.Render) ([]*Page, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to find the views path: %w", err)
	}
	dir := path.Join(wd, pagesPath)

	var pg []*Page
	err = filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(file) == ".gohtml" {
			pg = append(pg, newPage(render, file, dir))
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan pages: %w", err)
	}
	return pg, nil
}

func (p *Page) Render(w http.ResponseWriter, _ *http.Request) {
	err := p.render.HTML(w, http.StatusOK, p.Name, nil)
	if err != nil {
		log.Fatal(err)
	}
}

package render

import "github.com/thedevsaddam/renderer"

func New() *renderer.Render {
	return renderer.New(renderer.Options{ParseGlobPattern: "web/template/*.gohtml"})
}

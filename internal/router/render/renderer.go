package render

import "github.com/thedevsaddam/renderer"

const templatesGlobPattern = "web/template/*.gohtml"

func New() *renderer.Render {
	return renderer.New(renderer.Options{ParseGlobPattern: templatesGlobPattern})
}

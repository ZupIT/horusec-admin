package handler

import (
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type DefaultRender struct {
	render *renderer.Render
}

func NewDefaultRender(render *renderer.Render) *DefaultRender {
	return &DefaultRender{render: render}
}

func (h *DefaultRender) HandlerFunc(template string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.render.HTML(w, http.StatusOK, template, nil)
		if err != nil {
			panic(err)
		}
	}
}

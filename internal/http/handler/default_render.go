package handler

import (
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
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
			log.Fatal(err)
		}
	}
}

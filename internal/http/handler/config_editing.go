package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type ConfigEditing struct {
	render *renderer.Render
}

func NewConfigEditing(render *renderer.Render) *ConfigEditing {
	return &ConfigEditing{render: render}
}

func (h *ConfigEditing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.render.JSON(w, http.StatusOK, "{}")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to render JSON: %w", err))
	}
}

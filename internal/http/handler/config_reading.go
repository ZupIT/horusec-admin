package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type ConfigReading struct {
	render *renderer.Render
}

func NewConfigReading(render *renderer.Render) *ConfigReading {
	return &ConfigReading{render: render}
}

func (h *ConfigReading) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.render.JSON(w, http.StatusOK, "{}")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to render JSON: %w", err))
	}
}

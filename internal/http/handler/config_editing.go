package handler

import (
	"net/http"

	"github.com/tiagoangelozup/horusec-admin/internal/logger"

	"github.com/thedevsaddam/renderer"
)

type ConfigEditing struct {
	render *renderer.Render
}

func NewConfigEditing(render *renderer.Render) *ConfigEditing {
	return &ConfigEditing{render: render}
}

func (h *ConfigEditing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.WithPrefix("handler")

	err := h.render.JSON(w, http.StatusOK, "{}")
	if err != nil {
		panic(err)
	}

	log.Debug("the configuration has been successfully edited")
}

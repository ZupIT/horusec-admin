package handler

import (
	"net/http"

	"github.com/tiagoangelozup/horusec-admin/internal/logger"

	"github.com/thedevsaddam/renderer"
)

type ConfigReading struct {
	render *renderer.Render
}

func NewConfigReading(render *renderer.Render) *ConfigReading {
	return &ConfigReading{render: render}
}

func (h *ConfigReading) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.WithPrefix("handler")

	err := h.render.JSON(w, http.StatusOK, "{}")
	if err != nil {
		panic(err)
	}

	log.Debug("the configuration was found")
}

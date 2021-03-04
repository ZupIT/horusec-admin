package handler

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/core"

	"github.com/thedevsaddam/renderer"
)

type (
	ConfigReading struct {
		render *renderer.Render
		reader ConfigReader
	}
	ConfigReader interface {
		GetConfig() (*core.Configuration, error)
	}
)

func NewConfigReading(render *renderer.Render, reader ConfigReader) *ConfigReading {
	return &ConfigReading{render: render, reader: reader}
}

func (h *ConfigReading) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cfg, err := h.reader.GetConfig()
	if err != nil {
		panic(err)
	}

	// Answer
	if err = h.render.JSON(w, http.StatusOK, cfg); err != nil {
		panic(err)
	}
}

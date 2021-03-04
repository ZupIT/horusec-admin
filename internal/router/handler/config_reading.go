package handler

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/pkg/core"

	"github.com/thedevsaddam/renderer"
)

type (
	ConfigReading struct {
		render *renderer.Render
		reader core.ConfigurationReader
	}
)

func NewConfigReading(render *renderer.Render, reader core.ConfigurationReader) *ConfigReading {
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

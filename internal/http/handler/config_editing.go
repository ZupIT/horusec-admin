package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/core"

	"github.com/thedevsaddam/renderer"
)

type (
	ConfigEditing struct {
		render *renderer.Render
		writer ConfigWriter
	}
	ConfigWriter interface {
		Update(*core.Configuration) error
	}
)

func NewConfigEditing(render *renderer.Render, writer ConfigWriter) *ConfigEditing {
	return &ConfigEditing{render: render, writer: writer}
}

func (h *ConfigEditing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Unmarshall request body
	cfg := new(core.Configuration)
	if err := json.NewDecoder(r.Body).Decode(cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update configurations
	err := h.writer.Update(cfg)
	if err != nil {
		panic(err)
	}

	// Answer
	w.WriteHeader(http.StatusNoContent)
}

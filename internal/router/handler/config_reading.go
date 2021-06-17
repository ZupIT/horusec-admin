// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/tracing"
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
	span, ctx := tracing.StartSpanFromContext(r.Context(), "internal/router/handler.(*ConfigReading).ServeHTTP")
	defer span.Finish()

	cfg, err := h.reader.GetConfig(ctx)
	if err != nil {
		span.SetError(err)
		panic(err)
	}

	response := map[string]interface{}{
		"spec": cfg.Spec,
		"status": cfg.Status,
	}

	// Answer
	if err = h.render.JSON(w, http.StatusOK, response); err != nil {
		span.SetError(err)
		panic(err)
	}
}

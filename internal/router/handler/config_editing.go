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
	"io/ioutil"
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/logger"

	"github.com/thedevsaddam/renderer"

	"github.com/ZupIT/horusec-admin/internal/tracing"
	"github.com/ZupIT/horusec-admin/pkg/core"
)

type (
	ConfigEditing struct {
		render *renderer.Render
		writer core.ConfigurationWriter
	}
)

func NewConfigEditing(render *renderer.Render, writer core.ConfigurationWriter) *ConfigEditing {
	return &ConfigEditing{render: render, writer: writer}
}

//nolint:funlen // breaking the method 'ServeHTTP' is infeasible
func (h *ConfigEditing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracing.StartSpanFromContext(r.Context(), "internal/router/handler.(*ConfigEditing).ServeHTTP")
	defer span.Finish()

	// Unmarshall request body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		span.SetError(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update configurations
	if err := h.writer.CreateOrUpdate(ctx, payload); err != nil {
		span.SetError(err)
		panic(err)
	}

	// Answer
	w.WriteHeader(http.StatusNoContent)
	logger.WithPrefix(ctx, "[204]").Info("Configurations has been applied!")
}

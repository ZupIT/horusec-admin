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

package middleware

import (
	"context"
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/tracing"
	"github.com/go-chi/chi/middleware"
)

type TraceInitializer struct{}

func NewTracer() *TraceInitializer {
	return &TraceInitializer{}
}

func (t *TraceInitializer) Initialize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := tracing.StartSpanFromRequest(r)
		defer func() {
			defer span.Finish()
			if err := recover(); err != nil {
				span.Panic(err)
			}
		}()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		ctx = context.WithValue(ctx, middleware.RequestIDKey, span.String())
		next.ServeHTTP(ww, r.WithContext(ctx))
		span.SetHTTPResponseStatus(ww.Status())
	})
}

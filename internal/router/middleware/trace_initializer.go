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
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type TraceInitializer struct{}

func NewTracer() *TraceInitializer {
	return &TraceInitializer{}
}

func (t *TraceInitializer) Initialize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := t.startSpanFromRequest(r)
		ctx = context.WithValue(ctx, middleware.RequestIDKey, fmt.Sprintf("%+v", span))

		defer func() {
			defer span.Finish()
			if err := recover(); err != nil {
				ext.HTTPStatusCode.Set(span, uint16(500))
				ext.Error.Set(span, true)
				span.SetTag("error.type", "panic")
				span.LogKV("event", "error",
					"error.kind", "panic",
					"message", err,
					"stack", string(debug.Stack()))
				panic(err)
			}
		}()

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		ext.HTTPUrl.Set(span, fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI))
		ext.HTTPMethod.Set(span, r.Method)
		span.SetTag("http.protocol", r.Proto)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r.WithContext(ctx))

		status := ww.Status()
		ext.HTTPStatusCode.Set(span, uint16(status))

		if status >= 500 && status < 600 {
			ext.Error.Set(span, true)
			span.SetTag("error.type", fmt.Sprintf("%d: %s", status, http.StatusText(status)))
			span.LogKV(
				"event", "error",
				"message", fmt.Sprintf("%d: %s", status, http.StatusText(status)),
			)
		}
	})
}

func (t *TraceInitializer) startSpanFromRequest(r *http.Request) (opentracing.Span, context.Context) {
	tracer := opentracing.GlobalTracer()
	ctx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	operation := r.URL.Path
	operation = r.Method + " " + operation
	return opentracing.StartSpanFromContextWithTracer(r.Context(), tracer, operation, ext.RPCServerOption(ctx))
}

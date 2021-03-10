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

package tracing

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
)

// Initialize create an instance of Jaeger Tracer and sets it as GlobalTracer.
func Initialize(service string, logger Logger) (io.Closer, error) {
	cfg, err := (&config.Configuration{ServiceName: service}).FromEnv()
	if err != nil {
		return nil, fmt.Errorf("cannot init Jaeger: %w", err)
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(&defaultLogger{logger}))
	if err != nil {
		return nil, fmt.Errorf("cannot init Jaeger: %w", err)
	}

	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}

func StartSpanFromContext(ctx context.Context, operationName string) (*Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	return &Span{Span: span}, ctx
}

func StartSpanFromRequest(r *http.Request) (*Span, context.Context) {
	ctx := ExtractSpanContextFromRequest(r)
	span, ctxWithSpan := opentracing.StartSpanFromContext(r.Context(), r.Method+" "+r.URL.Path, ext.RPCServerOption(ctx))

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	ext.HTTPUrl.Set(span, fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI))
	ext.HTTPMethod.Set(span, r.Method)
	span.SetTag("http.protocol", r.Proto)

	return &Span{Span: span}, ctxWithSpan
}

func SpanFromContext(ctx context.Context) *Span {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}

	return &Span{Span: span}
}

func ExtractSpanContextFromRequest(r *http.Request) opentracing.SpanContext {
	tracer := opentracing.GlobalTracer()
	ctx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	return ctx
}

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

type Tracer struct {
	opentracing.Tracer
}

func NewTracer(tracer opentracing.Tracer) *Tracer {
	return &Tracer{Tracer: tracer}
}

func (t *Tracer) Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ctx := t.startSpanFromRequest(r)
		defer span.Finish()

		defer func() {
			if err := recover(); err != nil {
				ext.HTTPStatusCode.Set(span, uint16(500))
				ext.Error.Set(span, true)
				span.SetTag("error.type", "panic")
				span.LogKV("event", "error",
					"error.kind", "panic",
					"message", err,
					"stack", string(debug.Stack()))
				span.Finish()

				panic(err)
			}
		}()
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, r.URL.Path)

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

func (t *Tracer) startSpanFromRequest(r *http.Request) (opentracing.Span, context.Context) {
	ctx, _ := t.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	operation := r.URL.Path
	operation = r.Method + " " + operation
	return opentracing.StartSpanFromContextWithTracer(r.Context(), t, operation, ext.RPCServerOption(ctx))
}

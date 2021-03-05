package logger

import (
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

type (
	RequestFormatter struct {
		*middleware.DefaultLogFormatter
	}
	request struct {
		*log.Entry
	}
)

func NewRequestFormatter() *RequestFormatter {
	formatter := &middleware.DefaultLogFormatter{Logger: newRequest()}
	return &RequestFormatter{DefaultLogFormatter: formatter}
}

func newRequest() *request {
	return &request{Entry: WithPrefix("request")}
}

func (r *request) Print(v ...interface{}) {
	r.Entry.Trace(v...)
}

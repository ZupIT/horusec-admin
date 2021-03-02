package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ZupIT/horusec-admin/internal/logger"
)

const (
	Addr            = ":3000"
	ShutdownTimeout = 5 * time.Second
)

type (
	server struct {
		*http.Server
	}
	Interface interface {
		Start() Interface
		GracefullyShutdown() error
	}
)

func New(handler http.Handler) Interface {
	return &server{Server: &http.Server{Addr: Addr, Handler: handler}}
}

func (s *server) Start() Interface {
	go func() {
		log := logger.WithPrefix("server")
		log.WithField("addr", Addr).Info("listening")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Warn("listen error")
		}
	}()

	return s
}

func (s *server) GracefullyShutdown() error {
	logger.WithPrefix("server").Warn("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}

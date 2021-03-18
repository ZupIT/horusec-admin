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

package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	serverconfig "github.com/ZupIT/horusec-admin/config/server"
	"github.com/ZupIT/horusec-admin/internal/logger"
)

const (
	ShutdownTimeout = 5 * time.Second
)

type (
	server struct {
		*http.Server
		Config *serverconfig.Config
	}
	Interface interface {
		Start() Interface
		GracefullyShutdown() error
	}
)

func New(handler http.Handler, config *serverconfig.Config) Interface {
	return &server{
		Server: &http.Server{Addr: config.GetAddr(), Handler: handler},
		Config: config,
	}
}

func (s *server) Start() Interface {
	go func() {
		log := logger.WithPrefix(context.Background(), "server")
		log.WithField("addr", s.Server.Addr).Info("listening")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("listen error")
		}
	}()

	return s
}

func (s *server) GracefullyShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	logger.WithPrefix(ctx, "server").Warn("shutting down server")
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}

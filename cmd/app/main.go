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

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZupIT/horusec-admin/internal/logger"
	"github.com/ZupIT/horusec-admin/internal/server"
	"github.com/ZupIT/horusec-admin/internal/tracing"
)

var log = logger.WithPrefix(context.TODO(), "main")

// nolint:funlen // main will initialize all components
func main() {
	closer, err := tracing.Initialize("horusec-admin", logger.WithPrefix(context.TODO(), "tracing"))
	if err != nil {
		log.WithError(err).Fatal("failed to initialize tracer")
	}
	defer closer.Close()

	r, err := newRouter()
	if err != nil {
		log.WithError(err).Fatal("failed to create HTTP request router")
	}

	srv := server.New(r).Start()

	waitForInterruptSignal()
	if err = srv.GracefullyShutdown(); err != nil {
		log.WithError(err).Fatal("server forced to shutdown")
	}

	log.Info("server exiting")
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

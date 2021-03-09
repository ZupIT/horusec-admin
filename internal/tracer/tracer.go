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

package tracer

import (
	"fmt"
	"io"

	"github.com/ZupIT/horusec-admin/internal/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// New returns an instance of Jaeger Tracer.
func New(service string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := (&config.Configuration{ServiceName: service}).FromEnv()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot init Jaeger: %w", err)
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(logger.NewJaeger()))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot init Jaeger: %w", err)
	}

	return tracer, closer, nil
}

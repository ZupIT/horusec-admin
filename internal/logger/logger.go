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

package logger

import (
	"context"
	"os"
	"strings"

	"github.com/ZupIT/horusec-admin/internal/tracing"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// nolint
func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05 -0700",
		FullTimestamp:   true,
	})
	level := resolveLogLevelFromEnv()
	log.SetLevel(level)
}

func WithPrefix(ctx context.Context, prefix string) *log.Entry {
	entry := log.WithField("prefix", prefix).WithContext(ctx)
	if span := tracing.SpanFromContext(ctx); span != nil {
		if info := span.Info(); info != nil {
			entry = entry.WithFields(log.Fields{
				"trace":   info.TraceID,
				"span":    info.SpanID,
				"parent":  info.ParentID,
				"sampled": info.IsSampled,
			})
		}
	}
	return entry
}

func IsTrace() bool {
	return log.GetLevel() == log.TraceLevel
}

func resolveLogLevelFromEnv() log.Level {
	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		return log.InfoLevel
	}

	l, err := log.ParseLevel(strings.ToLower(level))
	if err != nil {
		log.Warnf("provided LOG_LEVEL %s is invalid. Fallback to info.", os.Getenv("LOG_LEVEL"))
		return log.InfoLevel
	}
	return l
}

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
	return &request{Entry: WithPrefix(context.Background(), "request")}
}

func (r *request) Print(v ...interface{}) {
	r.Entry.Trace(v...)
}

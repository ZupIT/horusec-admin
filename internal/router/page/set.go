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

package page

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/router/handler"
)

type (
	// page describes a single HTML page route.
	page struct {
		Pattern string
		Handler http.Handler
	}
	// Set describes all HTML pages routes.
	Set []*page
)

// NewSet creates and returns all HTML pages routes.
func NewSet(defaultRender *handler.DefaultRender) Set {
	return []*page{
		{Pattern: "/", Handler: defaultRender.HandlerFunc("index")},
		{Pattern: "/config-auth", Handler: defaultRender.HandlerFunc("config-auth")},
		{Pattern: "/config-general", Handler: defaultRender.HandlerFunc("config-general")},
		{Pattern: "/config-hosts", Handler: defaultRender.HandlerFunc("config-hosts")},
		{Pattern: "/config-resources", Handler: defaultRender.HandlerFunc("config-resources")},
		{Pattern: "/home", Handler: defaultRender.HandlerFunc("home")},
		{Pattern: "/not-authorized", Handler: defaultRender.HandlerFunc("not-authorized")},
	}
}

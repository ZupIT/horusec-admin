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

package api

import (
	"net/http"
)

type (
	// api describes a single Rest API route.
	api struct {
		Name          string
		Method        string
		Pattern       string
		Handler       http.Handler
		Authenticated bool
	}
	// Set describes all Rest API routes.
	Set []*api
)

// NewSet creates and returns all Rest API routes.
// nolint:funlen // it contains all definitions of all APIs and for that reason this list can be as long as needed
func NewSet(h *Handlers) Set {
	return []*api{
		{
			Name:          "Configurations",
			Method:        "GET",
			Pattern:       "/config",
			Handler:       h.ConfigReading,
			Authenticated: true,
		},
		{
			Name:          "Configurations",
			Method:        "PATCH",
			Pattern:       "/config",
			Handler:       h.ConfigEditing,
			Authenticated: true,
		},
		{
			Name:          "HealthCheck",
			Method:        "GET",
			Pattern:       "/health",
			Handler:       h.Health,
			Authenticated: true,
		},
		{
			Name:          "Authentication",
			Method:        "POST",
			Pattern:       "/auth",
			Handler:       h.Auth,
			Authenticated: true,
		},
	}
}

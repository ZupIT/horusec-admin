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
func NewSet(h *Handlers) Set {
	return []*api{
		{
			Name:          "Configurations",
			Method:        "GET",
			Pattern:       "/config",
			Handler:       h.ConfigReading,
			Authenticated: false,
		},
		{
			Name:          "Configurations",
			Method:        "PATCH",
			Pattern:       "/config",
			Handler:       h.ConfigEditing,
			Authenticated: false,
		},
		{
			Name:          "HealthCheck",
			Method:        "GET",
			Pattern:       "/health",
			Handler:       h.Health,
			Authenticated: false,
		},
		{
			Name:          "Authentication",
			Method:        "POST",
			Pattern:       "/auth",
			Handler:       h.Auth,
			Authenticated: false,
		},
	}
}

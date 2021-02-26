package router

import (
	"net/http"
)

// API describes a single Rest API route.
type API struct {
	Name          string
	Method        string
	Pattern       string
	Handler       http.Handler
	Authenticated bool
}

// newAPIs creates and returns all Rest API routes.
func newAPIs(h *handlers) []*API {
	return []*API{
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

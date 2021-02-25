package router

import (
	"net/http"
)

// Route describes a single route.
type Route struct {
	Name          string
	Method        string
	Pattern       string
	Handler       http.Handler
	Authenticated bool
}

// newRoutes creates and returns all the API routes.
func newRoutes(h *handlers) []*Route {
	return []*Route{
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

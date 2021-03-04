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
	// Set describes all HTML pages routes
	Set []*page
)

// NewSet creates and returns all HTML pages routes.
func NewSet(defaultRender *handler.DefaultRender) Set {
	return []*page{
		{Pattern: "/", Handler: defaultRender.HandlerFunc("index")},
		{Pattern: "/config-auth", Handler: defaultRender.HandlerFunc("config-auth")},
		{Pattern: "/config-general", Handler: defaultRender.HandlerFunc("config-general")},
		{Pattern: "/config-manager", Handler: defaultRender.HandlerFunc("config-manager")},
		{Pattern: "/home", Handler: defaultRender.HandlerFunc("home")},
		{Pattern: "/not-authorized", Handler: defaultRender.HandlerFunc("not-authorized")},
	}
}

package router

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/http/handler"
)

type (
	// Page describes a single HTML page route.
	Page struct {
		Pattern string
		Handler http.Handler
	}
)

// newPages creates and returns all HTML pages routes.
func newPages(defaultRender *handler.DefaultRender) []*Page {
	return []*Page{
		{Pattern: "/", Handler: defaultRender.HandlerFunc("index")},
		{Pattern: "/config-auth", Handler: defaultRender.HandlerFunc("config-auth")},
		{Pattern: "/config-general", Handler: defaultRender.HandlerFunc("config-general")},
		{Pattern: "/config-manager", Handler: defaultRender.HandlerFunc("config-manager")},
		{Pattern: "/home", Handler: defaultRender.HandlerFunc("home")},
		{Pattern: "/not-authorized", Handler: defaultRender.HandlerFunc("not-authorized")},
	}
}

package router

import "github.com/tiagoangelozup/horusec-admin/internal/http/handler"

type handlers struct {
	Auth          *handler.Auth
	ConfigEditing *handler.ConfigEditing
	ConfigReading *handler.ConfigReading
	Health        *handler.Health
}

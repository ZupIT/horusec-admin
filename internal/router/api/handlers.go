package api

import "github.com/ZupIT/horusec-admin/internal/router/handler"

type Handlers struct {
	Auth          *handler.Auth
	ConfigEditing *handler.ConfigEditing
	ConfigReading *handler.ConfigReading
	Health        *handler.Health
}

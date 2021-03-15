package kubernetes

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewHorusecManagerClient,
	NewRestConfig,
	wire.Struct(new(ObjectComparator), "*"),
)

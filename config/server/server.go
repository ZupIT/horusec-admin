package serverconfig

import (
	"fmt"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
)

type Config struct {
	addr string
}

const DefaultPort = 3000

func New() *Config {
	return &Config{
		addr: fmt.Sprintf(`:%b`, env.GetEnvOrDefaultInt("PORT", DefaultPort)),
	}
}

func (c *Config) GetAddr() string {
	return c.addr
}

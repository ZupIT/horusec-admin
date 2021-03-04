package core

type (
	ConfigurationWriter interface {
		Update(*Configuration) error
	}
	ConfigurationReader interface {
		GetConfig() (*Configuration, error)
	}
)

package core

type (
	ConfigurationWriter interface {
		CreateOrUpdate(*Configuration) error
	}
	ConfigurationReader interface {
		GetConfig() (*Configuration, error)
	}
)

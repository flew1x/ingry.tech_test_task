package config

const (
	restPath = "rest"
)

type IRESTConfig interface {
	GetHost() string
	GetPort() int
}

type RestConfig struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}

func NewRestConfig() *RestConfig {
	restConfig := &RestConfig{}
	mustUnmarshalStruct(restPath, &restConfig)

	return restConfig
}

func (c *RestConfig) GetHost() string {
	return c.Host
}

func (c *RestConfig) GetPort() int {
	return c.Port
}

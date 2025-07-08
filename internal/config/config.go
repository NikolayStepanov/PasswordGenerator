package config

const (
	defaultHttpHost = ""
	defaultHttpPort = "8080"
)

type Paths struct {
	Index string
}

type (
	Config struct {
		HTTP        HTTPConfig
		PathHandles Paths
	}

	HTTPConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
)

func Init() *Config {
	cfg := Config{}

	cfg.HTTP.Port = defaultHttpPort
	cfg.HTTP.Host = defaultHttpHost

	cfg.PathHandles = Paths{
		Index: "/",
	}

	return &cfg
}

// Package config provides application configuration structures and initialization logic
package config

import (
	"net/http"
	"time"
)

const (
	defaultHTTPHost = ""     // defaultHttpHost address for HTTP server to bind
	defaultHTTPPort = "8080" // default port for running the HTTP server
)

// Paths holds route paths for various HTTP handlers
type Paths struct {
	Password string
}

type (
	// Config holds the main application configuration,
	// including HTTP server settings and route paths
	Config struct {
		HTTP        HTTPConfig
		PathHandles Paths
	}
	// HTTPConfig contains configuration settings for the HTTP server.
	HTTPConfig struct {
		Host              string        `mapstructure:"host"`
		Port              string        `mapstructure:"port"`
		ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
		ReadTimeout       time.Duration `mapstructure:"read_timeout"`
		WriteTimeout      time.Duration `mapstructure:"write_timeout"`
		CommandTimeout    time.Duration `mapstructure:"command_timeout"`
		IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
	}
)

// Init initializes and returns a Config instance with default HTTP settings and path handles
func Init() *Config {
	cfg := Config{}

	cfg.HTTP.Port = defaultHTTPPort
	cfg.HTTP.Host = defaultHTTPHost
	cfg.HTTP.ReadHeaderTimeout = 5 * time.Second
	cfg.HTTP.ReadTimeout = 10 * time.Second
	cfg.HTTP.WriteTimeout = 10 * time.Second
	cfg.HTTP.IdleTimeout = 120 * time.Second

	cfg.PathHandles = Paths{
		Password: http.MethodPost + " " + "/password",
	}

	return &cfg
}

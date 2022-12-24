package config

import (
	"os"
	"strconv"
)

type envGetter interface {
	Getenv(key string) string
}

type cfg struct {
	envGetter envGetter
}

type osCfg struct{}

func (osCfg) Getenv(key string) string {
	return os.Getenv(key)
}

func New() *cfg {
	return &cfg{osCfg{}}
}

type Config struct {
	Server Server
}

type Server struct {
	Hostname string
	Port     int
}

const (
	cHostname = "HOSTNAME"
	cPort     = "PORT"
)

func (cfg *cfg) All() Config {
	return Config{
		Server: Server{
			Hostname: cfg.envString(cHostname, ""),
			Port:     cfg.envInt(cPort, 1323),
		},
	}
}

func (cfg *cfg) SetEnvGetter(overrideEnvGetter envGetter) {
	cfg.envGetter = overrideEnvGetter
}

func (cfg *cfg) envString(key, defaultValue string) string {
	value := cfg.envGetter.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (cfg *cfg) envInt(key string, defaultValue int) int {
	value := cfg.envGetter.Getenv(key)

	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

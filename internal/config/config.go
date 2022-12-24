package config

import "os"

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

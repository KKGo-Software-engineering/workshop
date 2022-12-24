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
	Server      Server
	FeatureFlag FeatureFlag
}

type Server struct {
	Hostname string
	Port     int
}

type FeatureFlag struct {
	IsLimitMaxSpend bool
}

const (
	cHostname = "HOSTNAME"
	cPort     = "PORT"

	cFlagIsLimitMaxSpend = "FLAG_IS_LIMIT_MAX_SPEND"
)

func (cfg *cfg) All() Config {
	return Config{
		Server: Server{
			Hostname: cfg.envString(cHostname, ""),
			Port:     cfg.envInt(cPort, 1323),
		},
		FeatureFlag: FeatureFlag{
			IsLimitMaxSpend: cfg.envBool(cFlagIsLimitMaxSpend, false),
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

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func (cfg *cfg) envBool(key string, defaultValue bool) bool {
	value := cfg.envGetter.Getenv(key)

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolValue
}

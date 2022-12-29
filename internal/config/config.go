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
	Server       Server
	FeatureFlag  FeatureFlag
	DBConnection string
}

type Server struct {
	Hostname string
	Port     int
}

type FeatureFlag struct {
	IsLimitMaxBalanceOnCreate bool `json:"isLimitMaxBalanceOnCreate"`
}

const (
	cHostname = "HOSTNAME"
	cPort     = "PORT"

	cFlagIsLimitMaxBalanceOnCreate = "FLAG_IS_LIMIT_MAX_SPEND_ON_CREATE"
	cDBConnection                  = "DB_CONNECTION"
)

const (
	dPort         = 1323
	dDBConnection = "postgresql://postgres:password@localhost:5432/banking?sslmode=disable"
)

func (cfg *cfg) All() Config {
	return Config{
		Server: Server{
			Hostname: cfg.envString(cHostname, ""),
			Port:     cfg.envInt(cPort, dPort),
		},
		FeatureFlag: FeatureFlag{
			IsLimitMaxBalanceOnCreate: cfg.envBool(cFlagIsLimitMaxBalanceOnCreate, false),
		},
		DBConnection: cfg.envString(cDBConnection, dDBConnection),
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

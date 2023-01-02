package config

import (
	"os"
	"strconv"
)

type env interface {
	Getenv(key string) string
}

type cfg struct {
	env env
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

func (cfg *cfg) SetEnvGetter(overrideEnvGetter env) {
	cfg.env = overrideEnvGetter
}

func (cfg *cfg) envString(key, defaultValue string) string {
	val := cfg.env.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (cfg *cfg) envInt(key string, defaultValue int) int {
	v := cfg.env.Getenv(key)

	val, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return val
}

func (cfg *cfg) envBool(key string, defaultValue bool) bool {
	v := cfg.env.Getenv(key)

	val, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return val
}

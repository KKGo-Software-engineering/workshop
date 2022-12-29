//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockOsCfg struct {
	mock.Mock
}

func (m *mockOsCfg) Getenv(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func TestStringEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue string
		osValue      string
		want         string
	}{
		{"get string from environment", "test-key", "default-value", "test-value", "test-value"},
		{"get defualt string if not config", "", "default-value", "", "default-value"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", tc.key).Return(tc.osValue)
			cfg.SetEnvGetter(mCfg)
			got := cfg.envString(tc.key, tc.defaultValue)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue int
		osValue      string
		want         int
	}{
		{"get int from environment", "test-key", 10, "1", 1},
		{"get defualt int if not config", "", 10, "", 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", tc.key).Return(tc.osValue)
			cfg.SetEnvGetter(mCfg)
			got := cfg.envInt(tc.key, tc.defaultValue)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBoolEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue bool
		osValue      string
		want         bool
	}{
		{"get bool from environment", "test-key", true, "FALSE", false},
		{"get defualt bool if not config", "", true, "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", tc.key).Return(tc.osValue)
			cfg.SetEnvGetter(mCfg)
			got := cfg.envBool(tc.key, tc.defaultValue)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetAllConfig(t *testing.T) {
	cfg := New()
	tests := []struct {
		name   string
		mCfgFn func() *mockOsCfg
		want   Config
	}{
		{"no config env should return default config",
			func() *mockOsCfg {
				return defaultMock(cHostname, cPort, cFlagIsLimitMaxBalanceOnCreate, cDBConnection)
			},
			Config{Server: Server{Port: 1323},
				DBConnection: dDBConnection}},
		{"config hostname env should return as changed", func() *mockOsCfg {
			mCfg := defaultMock(cPort, cFlagIsLimitMaxBalanceOnCreate, cDBConnection)
			mCfg.On("Getenv", cHostname).Return("test-hostname")
			return mCfg
		}, Config{Server: Server{Hostname: "test-hostname", Port: 1323},
			DBConnection: dDBConnection}},
		{"config port env should return as changed", func() *mockOsCfg {
			mCfg := defaultMock(cHostname, cFlagIsLimitMaxBalanceOnCreate, cDBConnection)
			mCfg.On("Getenv", cPort).Return("4444")
			return mCfg
		}, Config{Server: Server{Port: 4444},
			DBConnection: dDBConnection}},
		{"config bool FLAG_IS_LIMIT_MAX_SPEND_ON_CREATE env should return as changed", func() *mockOsCfg {
			mCfg := defaultMock(cPort, cHostname, cDBConnection)
			mCfg.On("Getenv", cFlagIsLimitMaxBalanceOnCreate).Return("TRUE")
			return mCfg
		}, Config{
			Server:       Server{Port: 1323},
			FeatureFlag:  FeatureFlag{IsLimitMaxBalanceOnCreate: true},
			DBConnection: dDBConnection}},
		{"config DB_CONNECTION env should return as changed", func() *mockOsCfg {
			mCfg := defaultMock(cPort, cHostname, cFlagIsLimitMaxBalanceOnCreate)
			mCfg.On("Getenv", cDBConnection).Return("test-db-connection")
			return mCfg
		}, Config{
			Server:       Server{Port: 1323},
			DBConnection: "test-db-connection"}},
		{"config all env should return as changed", func() *mockOsCfg {
			mCfg := defaultMock()
			mCfg.On("Getenv", cHostname).Return("test-hostname")
			mCfg.On("Getenv", cPort).Return("4444")
			mCfg.On("Getenv", cFlagIsLimitMaxBalanceOnCreate).Return("TRUE")
			mCfg.On("Getenv", cDBConnection).Return("test-db-connection")
			return mCfg
		}, Config{
			Server:       Server{Hostname: "test-hostname", Port: 4444},
			FeatureFlag:  FeatureFlag{IsLimitMaxBalanceOnCreate: true},
			DBConnection: "test-db-connection",
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := tc.mCfgFn()
			cfg.SetEnvGetter(mCfg)
			got := cfg.All()
			assert.Equal(t, tc.want, got)
		})
	}
}

func defaultMock(envs ...string) *mockOsCfg {
	mCfg := new(mockOsCfg)
	for _, env := range envs {
		mCfg.On("Getenv", env).Return("")
	}
	return mCfg
}

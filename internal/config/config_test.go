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
		{"no config env should return default config", func() *mockOsCfg {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", cHostname).Return("")
			mCfg.On("Getenv", cPort).Return("")
			return mCfg
		}, Config{Server: Server{Port: 1323}}},
		{"config hostname env should return as changed", func() *mockOsCfg {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", cHostname).Return("test-hostname")
			mCfg.On("Getenv", cPort).Return("")
			return mCfg
		}, Config{Server: Server{Hostname: "test-hostname", Port: 1323}}},
		{"config port env should return as changed", func() *mockOsCfg {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", cHostname).Return("")
			mCfg.On("Getenv", cPort).Return("4444")
			return mCfg
		}, Config{Server: Server{Hostname: "", Port: 4444}}},
		{"config hostname and port env should return as changed", func() *mockOsCfg {
			mCfg := new(mockOsCfg)
			mCfg.On("Getenv", cHostname).Return("test-hostname")
			mCfg.On("Getenv", cPort).Return("4444")
			return mCfg
		}, Config{Server: Server{Hostname: "test-hostname", Port: 4444}}},
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

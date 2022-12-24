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

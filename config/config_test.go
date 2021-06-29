package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigOK(t *testing.T) {
	cfg := New()
	err := cfg.GetFromFile("testdata/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	cases := []struct {
		expected string
		actual   string
	}{
		{
			expected: "debug",
			actual:   cfg.Logger.Level,
		},
		{
			expected: ":8080",
			actual:   cfg.Server.Addr,
		},
		{
			expected: "database URL",
			actual:   cfg.Storage.DatabaseURL,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.actual)
	}
}

func TestConfigFail(t *testing.T) {
	cfg := New()
	err := cfg.GetFromFile("testdata/apiserver_bad.toml")
	assert.NotEqual(t, nil, err)
}

func TestConfigNotExist(t *testing.T) {
	cfg := New()
	err := cfg.GetFromFile("testdata/notexisting.toml")
	assert.NotEqual(t, nil, err)
}

func TestConfigDefault(t *testing.T) {
	cfg := New()
	err := cfg.GetFromFile("testdata/apiserver_empty.toml")
	if err != nil {
		t.Error(err)
	}
	cases := []struct {
		expected string
		actual   string
	}{
		{
			expected: "debug",
			actual:   cfg.Logger.Level,
		},
		{
			expected: ":8082",
			actual:   cfg.Server.Addr,
		},
		{
			expected: "some default URL string",
			actual:   cfg.Storage.DatabaseURL,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.actual)
	}

}

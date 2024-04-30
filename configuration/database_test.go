package configuration_test

import (
	"github.com/PaluMacil/barkdognet/configuration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBConfig_ConnectionString(t *testing.T) {
	type testCase struct {
		name     string
		config   configuration.Database
		expected string
	}

	testCases := []testCase{
		{
			name: "database blank",
			config: configuration.Database{
				Host:     "10.10.10.10",
				Port:     "5432",
				User:     "user",
				Password: "password",
				Database: "",
			},
			expected: "postgres://user:password@10.10.10.10:5432/barkdog",
		},
		{
			name: "user is blank",
			config: configuration.Database{
				Host:     "localhost",
				Port:     "6262",
				User:     "",
				Password: "1212",
				Database: "database",
			},
			expected: "postgres://barkadmin:1212@localhost:6262/database",
		},
		{
			name: "password is blank",
			config: configuration.Database{
				Host:     "localhost",
				Port:     "6262",
				User:     "user",
				Password: "",
				Database: "database",
			},
			expected: "postgres://user@localhost:6262/database",
		},
		{
			name: "port is blank",
			config: configuration.Database{
				Host:     "localhost",
				Port:     "",
				User:     "user",
				Password: "password",
				Database: "database",
			},
			expected: "postgres://user:password@localhost:5432/database",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.config.ConnectionString())
		})
	}
}

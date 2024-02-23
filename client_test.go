package entx_test

import (
	"testing"

	"github.com/datumforge/entx"

	"github.com/stretchr/testify/assert"
)

func TestCheckDialect(t *testing.T) {
	testCases := []struct {
		name     string
		dialect  string
		expected string
		errorMsg string
	}{
		{
			name:     "sqlite",
			dialect:  "sqlite",
			expected: "sqlite3",
		},
		{
			name:     "libsql",
			dialect:  "sqlite",
			expected: "sqlite3",
		},
		{
			name:     "postgres",
			dialect:  "postgres",
			expected: "postgres",
		},
		{
			name:     "unsupported",
			dialect:  "mysql",
			errorMsg: "unsupported dialect: mysql",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			chk, err := entx.CheckEntDialect(tc.dialect)

			if tc.errorMsg == "" {
				assert.Nil(t, err)
				assert.NotNil(t, chk)
				assert.Equal(t, tc.expected, chk)
			} else {
				assert.NotNil(t, err)
				assert.Empty(t, chk)
				assert.ErrorContains(t, err, tc.errorMsg)
			}
		})
	}
}

func TestMultiWriteSupport(t *testing.T) {
	testCases := []struct {
		name     string
		dialect  string
		expected bool
	}{
		{
			name:     "sqlite",
			dialect:  "sqlite",
			expected: true,
		},
		{
			name:     "libsql",
			dialect:  "sqlite",
			expected: true,
		},
		{
			name:     "postgres",
			dialect:  "postgres",
			expected: false,
		},
		{
			name:     "unsupported",
			dialect:  "mysql",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			chk := entx.CheckMultiwriteSupport(tc.dialect)

			assert.Equal(t, tc.expected, chk)
		})
	}
}

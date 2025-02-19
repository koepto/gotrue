package models_test

import (
	"testing"

	"github.com/gobuffalo/pop/v5"
	"github.com/netlify/gotrue/models"
	"github.com/stretchr/testify/assert"
)

func TestTableNameNamespacing(t *testing.T) {
	cases := []struct {
		expected string
		value    interface{}
	}{
		{expected: "audit_log_entries", value: []*models.AuditLogEntry{}},
		{expected: "refresh_tokens", value: []*models.RefreshToken{}},
		{expected: "users", value: []*models.User{}},
	}

	for _, tc := range cases {
		m := &pop.Model{Value: tc.value}
		assert.Equal(t, tc.expected, m.TableName())
	}
}

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
		{expected: "audit_log_entry", value: []*models.AuditLogEntry{}},
		{expected: "instance", value: []*models.Instance{}},
		{expected: "refresh_token", value: []*models.RefreshToken{}},
		{expected: "\"user\"", value: []*models.User{}},
	}

	for _, tc := range cases {
		m := &pop.Model{Value: tc.value}
		assert.Equal(t, tc.expected, m.TableName())
	}
}

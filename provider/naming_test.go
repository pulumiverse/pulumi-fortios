package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalNameConversion(t *testing.T) {
	mod, name := convertName("fortios_firewall_securitypolicy_sort")
	assert.Equal(t, "firewall", mod)
	assert.Equal(t, "SecurityPolicySort", name)
}

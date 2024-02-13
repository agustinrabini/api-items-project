package tests

import (
	"testing"

	"github.com/agustinrabini/api-items-project/internal/platform/tests"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	response := executeRequest(buildRouter(), "GET", "/ping", nil, "")

	assert.Equal(t, "pong", response.Body.String())
}

package utils

import (
	"context"
	"testing"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/utils/tests"
)

// Deprecated: use tests.Context
func Context(t *testing.T) context.Context {
	return tests.Context(t)
}

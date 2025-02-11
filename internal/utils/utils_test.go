package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	require.True(t, GetValue(true))
}

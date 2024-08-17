package erkkratos

import (
	"testing"

	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNewErkFc(t *testing.T) {
	erk := NewErkFc(errors_example.ErrorServerDbError, "erk")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

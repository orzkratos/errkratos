package erkkratos

import (
	"testing"

	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNewEmtBottle_Wrap(t *testing.T) {
	emtBottle := NewEmtBottle(errors_example.ErrorServerDbError, "msg", "erk")
	erk := emtBottle.Wrap(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

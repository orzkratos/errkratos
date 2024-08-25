package erkkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNewErkFsc(t *testing.T) {
	erk := NewErkFsc(errors_example.ErrorServerDbError, "erk")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

func TestNewErkFmx(t *testing.T) {
	erk := NewErkFmx(errors_example.ErrorServerDbError, "msg")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

func TestNewErkFmx2(t *testing.T) {
	erk := NewErkFmx(errors_example.ErrorServerDbError, "msg")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)

	var err error = erk

	var target *errors.Error
	ok := errors.As(err, &target)
	t.Log(ok)
	t.Log(target)
}

func TestAs(t *testing.T) {
	{
		var erk = errors_example.ErrorServerDbError("wrong")
		var erx error = erk
		res, ok := As(erx)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	}

	{
		var erk *errors.Error
		var erx error = erk
		res, ok := As(erx)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	}
}

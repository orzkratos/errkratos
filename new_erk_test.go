package erkkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/erkkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNewErkFsC(t *testing.T) {
	erk := NewErkFsC(errors_example.ErrorServerDbError, "erk")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

func TestNewErkMtX(t *testing.T) {
	erk := NewErkMtX(errors_example.ErrorServerDbError, "msg")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)
}

func TestNewErkMtX_2(t *testing.T) {
	erk := NewErkMtX(errors_example.ErrorServerDbError, "msg")(erero.New("wac"))
	require.NotNil(t, erk)
	require.True(t, errors_example.IsServerDbError(erk))
	t.Log(erk)

	var erx error = erk

	var target *errors.Error
	ok := errors.As(erx, &target)
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

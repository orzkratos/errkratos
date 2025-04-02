package errkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestNewEro(t *testing.T) {
	ero := NewEro(errors_example.ErrorServerDbError("wrong reason=%s", "unknown"))
	require.NotNil(t, ero)
	require.NotNil(t, ero.Erk)

	t.Log(ero.Erk.String())
}

func TestFmtEro(t *testing.T) {
	ero := FmtEro(errors_example.ErrorServerDbError, "wrong reason=%s", "unknown")
	require.NotNil(t, ero)
	require.NotNil(t, ero.Erk)

	{
		var erx interface{} = ero
		err, ok := erx.(error) //不要实现 error 接口，而且注意一定不要实现，否则会加重开发者的心智负担
		require.False(t, ok)
		require.Nil(t, err)
	}

	{
		var erx interface{} = ero.Erk
		err, ok := erx.(error) //已经实现 error 接口
		require.True(t, ok)
		require.Error(t, err)
	}

	t.Log(ero.Erk.String())
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

func TestIs(t *testing.T) {
	erk1 := errors_example.ErrorServerDbError("wrong-1")
	erk2 := errors_example.ErrorServerDbError("wrong-2")
	require.True(t, Is(erk1, erk2))

	require.True(t, errors.Is(erk1, erk1)) //还是相等
	require.True(t, erero.Ise(erk1, erk1)) //依然相等
}

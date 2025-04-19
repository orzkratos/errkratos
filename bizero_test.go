package errkratos

import (
	"testing"

	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
)

func TestNewBizEro(t *testing.T) {
	ero := NewBizEro(errors_example.ErrorServerDbError("wrong reason=%s", "unknown"))
	require.NotNil(t, ero)
	require.NotNil(t, ero.Erk)

	t.Log(ero.Erk.String())
}

func TestNewBizEro_NotImplementErrorInterface(t *testing.T) {
	ero := NewBizEro(errors_example.ErrorServerDbError("wrong reason=%s", "unknown"))
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

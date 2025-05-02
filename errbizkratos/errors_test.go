package errbizkratos_test

import (
	"testing"

	"github.com/orzkratos/errkratos"
	"github.com/orzkratos/errkratos/errbizkratos"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
)

func TestNewEbz(t *testing.T) {
	ebz := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong reason=%s", "unknown"))
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	t.Log(ebz.Erk.String())
}

func TestNewEbz_NotImplementErrorInterface(t *testing.T) {
	ebz := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong reason=%s", "unknown"))
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	{
		var erx interface{} = ebz
		err, ok := erx.(error) //不要实现 error 接口，而且注意一定不要实现，否则会加重开发者的心智负担
		require.False(t, ok)
		require.Nil(t, err)
	}

	{
		var erx interface{} = ebz.Erk
		err, ok := erx.(error) //已经实现 error 接口
		require.True(t, ok)
		require.Error(t, err)
	}

	t.Log(ebz.Erk.String())
}

func TestFrom(t *testing.T) {
	ebz1 := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong"))
	var err error = ebz1.Erk
	ebz2 := errbizkratos.From(err)
	require.True(t, errkratos.Is(ebz1.Erk, ebz2.Erk))
}

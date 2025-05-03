package errbizkratos_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/errbizkratos"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
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

func TestAs(t *testing.T) {
	{
		var erk = errors_example.ErrorServerDbError("wrong")
		var err error = erk
		// t.Log(erk != nil) // true
		// t.Log(err != nil) // true

		res, ok := errbizkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	}

	{
		var erk *errors.Error
		var err error = erk
		// t.Log(erk != nil) // false
		// t.Log(err != nil) // true

		res, ok := errbizkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	}
}

func TestIs(t *testing.T) {
	ebz1 := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong-1"))
	ebz2 := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong-2"))
	require.True(t, errbizkratos.Is(ebz1, ebz2))

	require.True(t, errors.Is(ebz1.Erk, ebz1.Erk)) //还是相等
	require.True(t, erero.Ise(ebz1.Erk, ebz1.Erk)) //依然相等
}

func TestFrom(t *testing.T) {
	ebz1 := errbizkratos.NewEbz(errors_example.ErrorServerDbError("wrong"))
	var err error = ebz1.Erk
	ebz2 := errbizkratos.From(err)
	require.True(t, errbizkratos.Is(ebz1, ebz2))
}

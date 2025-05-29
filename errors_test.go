package errkratos_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestAs(t *testing.T) {
	{
		var erk = errors_example.ErrorServerDbError("wrong")
		var err error = erk
		// t.Log(erk != nil) // true
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := errkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	}

	{
		var erk *errors.Error
		var err error = erk
		// t.Log(erk != nil) // false
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := errkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	}
}

func TestIs(t *testing.T) {
	erk1 := errors_example.ErrorServerDbError("wrong-1")
	erk2 := errors_example.ErrorServerDbError("wrong-2")
	require.True(t, errkratos.Is(erk1, erk2))

	require.True(t, errors.Is(erk1, erk1)) //还是相等
	require.True(t, erero.Ise(erk1, erk1)) //依然相等
}

func TestFrom(t *testing.T) {
	erk1 := errors_example.ErrorServerDbError("wrong")
	var err error = erk1
	erk2 := errkratos.From(err)
	require.True(t, errkratos.Is(erk1, erk2))
}

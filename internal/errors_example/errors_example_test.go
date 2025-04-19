package errors_example_test

import (
	"testing"

	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestIsServerDbError(t *testing.T) {
	erk := errors_example.ErrorServerDbError("error=%v", erero.New("abc"))
	res := errors_example.IsServerDbError(erk)
	require.True(t, res)
}

func TestIsServerDbTransactionError(t *testing.T) {
	erk := errors_example.ErrorServerDbTransactionError("error=%v", erero.New("txn"))
	res := errors_example.IsServerDbTransactionError(erk)
	require.True(t, res)
}

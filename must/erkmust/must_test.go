package erkmust_test

import (
	"testing"

	"github.com/orzkratos/errkratos"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/orzkratos/errkratos/internal/tests"
	"github.com/orzkratos/errkratos/must/erkmust"
	"github.com/yyle88/must"
)

func TestDone(t *testing.T) {
	{
		var erk *errkratos.Erk
		erkmust.Done(erk)
	}

	tests.ExpectPanic(t, func() {
		var erk = errorspb.ErrorServerDbError("wrong db")
		erkmust.Done(erk)
	})

	tests.ExpectPanic(t, func() {
		var erk *errkratos.Erk
		must.Done(erk) //这里也不知道是什么原因，这个是不可用的，请用前面的判定函数
	})
}

func TestMust(t *testing.T) {
	{
		var erk *errkratos.Erk
		erkmust.Must(erk)
	}

	tests.ExpectPanic(t, func() {
		var erk = errorspb.ErrorServerDbTransactionError("wrong tx")
		erkmust.Must(erk)
	})

	tests.ExpectPanic(t, func() {
		var erk *errkratos.Erk
		must.Must(erk) //这里也不知道是什么原因，这个是不可用的，请用前面的判定函数
	})
}

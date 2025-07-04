package ebzkratosmust_test

import (
	"testing"

	"github.com/orzkratos/errkratos/ebzkratos"
	"github.com/orzkratos/errkratos/ebzkratos/ebzkratosmust"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/orzkratos/errkratos/internal/tests"
)

func TestDone(t *testing.T) {
	{
		var ebz *ebzkratos.Ebz
		ebzkratosmust.Done(ebz)
	}

	tests.ExpectPanic(t, func() {
		var erk = errors_example.ErrorServerDbError("wrong db")
		ebzkratosmust.Done(ebzkratos.New(erk))
	})

	// 直接编译不过
	// tests.ExpectPanic(t, func() {
	// 	var ebz *ebzkratos.Ebz
	// 	must.Done(ebz)
	// })
}

func TestMust(t *testing.T) {
	{
		var ebz *ebzkratos.Ebz
		ebzkratosmust.Must(ebz)
	}

	tests.ExpectPanic(t, func() {
		var ebz = errors_example.ErrorServerDbTransactionError("wrong tx")
		ebzkratosmust.Must(ebzkratos.New(ebz))
	})

	// 直接编译不过
	// tests.ExpectPanic(t, func() {
	// 	var ebz *ebzkratos.Ebz
	// 	must.Must(ebz)
	// })
}

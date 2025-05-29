package erkbzmust_test

import (
	"testing"

	"github.com/orzkratos/errkratos/ebzkratos"
	"github.com/orzkratos/errkratos/ebzkratos/erkbzmust"
	"github.com/orzkratos/errkratos/internal/errors_example"
	"github.com/orzkratos/errkratos/internal/tests"
)

func TestDone(t *testing.T) {
	{
		var ebz *ebzkratos.Ebz
		erkbzmust.Done(ebz)
	}

	tests.ExpectPanic(t, func() {
		var erk = errors_example.ErrorServerDbError("wrong db")
		erkbzmust.Done(ebzkratos.New(erk))
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
		erkbzmust.Must(ebz)
	}

	tests.ExpectPanic(t, func() {
		var ebz = errors_example.ErrorServerDbTransactionError("wrong tx")
		erkbzmust.Must(ebzkratos.New(ebz))
	})

	// 直接编译不过
	// tests.ExpectPanic(t, func() {
	// 	var ebz *ebzkratos.Ebz
	// 	must.Must(ebz)
	// })
}

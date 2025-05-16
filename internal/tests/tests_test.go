package tests_test

import (
	"testing"

	"github.com/orzkratos/errkratos/internal/tests"
	"github.com/pkg/errors"
)

func TestExpectPanic(t *testing.T) {
	tests.ExpectPanic(t, func() {
		panic(errors.New("expect-panic"))
	})
}

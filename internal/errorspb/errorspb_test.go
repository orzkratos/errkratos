package errorspb_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestErrorUnknown(t *testing.T) {
	erk := errorspb.ErrorUnknown("internal system failure: %s", "database connection lost")
	t.Log(erk)
	t.Log(neatjsons.S(erk))

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_UNKNOWN.String(), erk.Reason)
	require.Equal(t, "internal system failure: database connection lost", erk.Message)
}

func TestErrorServerDbError(t *testing.T) {
	erk := errorspb.ErrorServerDbError("database query failed: %s", "connection timeout")
	t.Log(erk)
	t.Log(neatjsons.S(erk))

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_SERVER_DB_ERROR.String(), erk.Reason)
	require.Equal(t, "database query failed: connection timeout", erk.Message)
}

func TestErrorServerDbTransactionError(t *testing.T) {
	erk := errorspb.ErrorServerDbTransactionError("transaction failed: %s", "deadlock detected")
	t.Log(erk)
	t.Log(neatjsons.S(erk))

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_SERVER_DB_TRANSACTION_ERROR.String(), erk.Reason)
	require.Equal(t, "transaction failed: deadlock detected", erk.Message)
}

func TestIsUnknown(t *testing.T) {
	t.Run("match unknown", func(t *testing.T) {
		erk := errorspb.ErrorUnknown("test unknown error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsUnknown(erk))
	})

	t.Run("not match db", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("test db error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.False(t, errorspb.IsUnknown(erk))
	})

	t.Run("match standard", func(t *testing.T) {
		erk := errors.New(500, "UNKNOWN", "standard unknown error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsUnknown(erk))
	})

	t.Run("nil check", func(t *testing.T) {
		require.False(t, errorspb.IsUnknown(nil))
	})
}

func TestIsServerDbError(t *testing.T) {
	t.Run("match db", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("test db error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsServerDbError(erk))
	})

	t.Run("not match unknown", func(t *testing.T) {
		erk := errorspb.ErrorUnknown("test unknown")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.False(t, errorspb.IsServerDbError(erk))
	})

	t.Run("match standard", func(t *testing.T) {
		erk := errors.New(500, "SERVER_DB_ERROR", "standard db error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsServerDbError(erk))
	})
}

func TestIsServerDbTransactionError(t *testing.T) {
	t.Run("match transaction", func(t *testing.T) {
		erk := errorspb.ErrorServerDbTransactionError("test transaction error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsServerDbTransactionError(erk))
	})

	t.Run("not match db", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("test db error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.False(t, errorspb.IsServerDbTransactionError(erk))
	})

	t.Run("match standard", func(t *testing.T) {
		erk := errors.New(500, "SERVER_DB_TRANSACTION_ERROR", "standard transaction error")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.True(t, errorspb.IsServerDbTransactionError(erk))
	})
}

func TestErrorReasonEnumValues(t *testing.T) {
	// Test enum constant values match expected numbers
	require.Equal(t, int32(0), int32(errorspb.ErrorReason_UNKNOWN))
	require.Equal(t, int32(50001), int32(errorspb.ErrorReason_SERVER_DB_ERROR))
	require.Equal(t, int32(50002), int32(errorspb.ErrorReason_SERVER_DB_TRANSACTION_ERROR))
}

func TestErrorReasonStringConversion(t *testing.T) {
	// Test string representation
	require.Equal(t, "UNKNOWN", errorspb.ErrorReason_UNKNOWN.String())
	require.Equal(t, "SERVER_DB_ERROR", errorspb.ErrorReason_SERVER_DB_ERROR.String())
	require.Equal(t, "SERVER_DB_TRANSACTION_ERROR", errorspb.ErrorReason_SERVER_DB_TRANSACTION_ERROR.String())
}

func TestAllErrorTypesHaveCorrectHttpCode(t *testing.T) {
	t.Run("unknown code", func(t *testing.T) {
		erk := errorspb.ErrorUnknown("test")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.Equal(t, int32(500), erk.Code)
	})

	t.Run("db code", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("test")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.Equal(t, int32(500), erk.Code)
	})

	t.Run("transaction code", func(t *testing.T) {
		erk := errorspb.ErrorServerDbTransactionError("test")
		t.Log(erk)
		t.Log(neatjsons.S(erk))
		require.Equal(t, int32(500), erk.Code)
	})
}

func TestCrossErrorTypeChecking(t *testing.T) {
	t.Run("db check", func(t *testing.T) {
		erk := errorspb.ErrorServerDbError("database connection failed")
		t.Log(erk)
		t.Log(neatjsons.S(erk))

		require.True(t, errorspb.IsServerDbError(erk))
		require.False(t, errorspb.IsUnknown(erk))
		require.False(t, errorspb.IsServerDbTransactionError(erk))
	})

	t.Run("transaction check", func(t *testing.T) {
		erk := errorspb.ErrorServerDbTransactionError("transaction rollback failed")
		t.Log(erk)
		t.Log(neatjsons.S(erk))

		require.True(t, errorspb.IsServerDbTransactionError(erk))
		require.False(t, errorspb.IsUnknown(erk))
		require.False(t, errorspb.IsServerDbError(erk))
	})
}

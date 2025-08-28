package errorspb_test

import (
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errgenkratos"
	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	errgenkratos.SetMetadataFieldName("numeric_error_code")
	m.Run()
}

func TestErrorUnknown(t *testing.T) {
	erk := errorspb.ErrorUnknown("internal system failure: %s", "database connection lost")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_UNKNOWN.String(), erk.Reason)
	require.Equal(t, "internal system failure: database connection lost", erk.Message)
}

func TestErrorServerDbError(t *testing.T) {
	erk := errorspb.ErrorServerDbError("database query failed: %s", "connection timeout")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_SERVER_DB_ERROR.String(), erk.Reason)
	require.Equal(t, "database query failed: connection timeout", erk.Message)
}

func TestErrorServerDbTransactionError(t *testing.T) {
	erk := errorspb.ErrorServerDbTransactionError("transaction failed: %s", "deadlock detected")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, errorspb.ErrorReason_SERVER_DB_TRANSACTION_ERROR.String(), erk.Reason)
	require.Equal(t, "transaction failed: deadlock detected", erk.Message)
}

func TestIsUnknown(t *testing.T) {
	erk := errorspb.ErrorUnknown("test unknown error")
	t.Log(erk)
	require.True(t, errorspb.IsUnknown(erk))

	dbErk := errorspb.ErrorServerDbError("test db error")
	t.Log(dbErk)
	require.False(t, errorspb.IsUnknown(dbErk))

	stdErk := errors.New(500, "UNKNOWN", "standard unknown error")
	t.Log(stdErk)
	require.True(t, errorspb.IsUnknown(stdErk))

	require.False(t, errorspb.IsUnknown(nil))
}

func TestIsServerDbError(t *testing.T) {
	erk := errorspb.ErrorServerDbError("test db error")
	t.Log(erk)
	require.True(t, errorspb.IsServerDbError(erk))

	unknownErk := errorspb.ErrorUnknown("test unknown")
	t.Log(unknownErk)
	require.False(t, errorspb.IsServerDbError(unknownErk))

	stdErk := errors.New(500, "SERVER_DB_ERROR", "standard db error")
	t.Log(stdErk)
	require.True(t, errorspb.IsServerDbError(stdErk))
}

func TestIsServerDbTransactionError(t *testing.T) {
	erk := errorspb.ErrorServerDbTransactionError("test transaction error")
	t.Log(erk)
	require.True(t, errorspb.IsServerDbTransactionError(erk))

	dbErk := errorspb.ErrorServerDbError("test db error")
	t.Log(dbErk)
	require.False(t, errorspb.IsServerDbTransactionError(dbErk))

	stdErk := errors.New(500, "SERVER_DB_TRANSACTION_ERROR", "standard transaction error")
	t.Log(stdErk)
	require.True(t, errorspb.IsServerDbTransactionError(stdErk))
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

func TestMetadataFieldName(t *testing.T) {
	erk := errorspb.ErrorServerDbError("test with metadata")
	t.Log(erk)

	require.NotNil(t, erk.Metadata)

	internalCode, exists := erk.Metadata["numeric_error_code"]
	require.True(t, exists)

	expectedCode := fmt.Sprintf("%d", errorspb.ErrorReason_SERVER_DB_ERROR.Number())
	require.Equal(t, expectedCode, internalCode)
}

func TestAllErrorTypesHaveCorrectHttpCode(t *testing.T) {
	// All errors should have HTTP 500 status code
	unknownErk := errorspb.ErrorUnknown("test")
	t.Log(unknownErk)
	require.Equal(t, int32(500), unknownErk.Code)

	dbErk := errorspb.ErrorServerDbError("test")
	t.Log(dbErk)
	require.Equal(t, int32(500), dbErk.Code)

	txErk := errorspb.ErrorServerDbTransactionError("test")
	t.Log(txErk)
	require.Equal(t, int32(500), txErk.Code)
}

func TestCrossErrorTypeChecking(t *testing.T) {
	// Create one error and test it against all Is functions
	dbErk := errorspb.ErrorServerDbError("database connection failed")
	t.Log(dbErk)

	// Should only match its own type
	require.True(t, errorspb.IsServerDbError(dbErk))
	require.False(t, errorspb.IsUnknown(dbErk))
	require.False(t, errorspb.IsServerDbTransactionError(dbErk))

	// Test with transaction error
	txErk := errorspb.ErrorServerDbTransactionError("transaction rollback failed")
	t.Log(txErk)

	require.True(t, errorspb.IsServerDbTransactionError(txErk))
	require.False(t, errorspb.IsUnknown(txErk))
	require.False(t, errorspb.IsServerDbError(txErk))
}

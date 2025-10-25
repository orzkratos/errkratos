package newerk_test

import (
	"os"
	"testing"

	"github.com/orzkratos/errkratos/internal/errorspb"
	"github.com/orzkratos/errkratos/newerk"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Set reason code field name once to avoid concurrent modification
	// 设置一次原因码字段名，避免并发修改
	newerk.SetReasonCodeFieldName("numeric_reason_code_enum")
	os.Exit(m.Run())
}

func TestGetReasonCodeFieldName(t *testing.T) {
	// Verify reason code field name is set correctly
	// 验证原因码字段名设置正确
	require.Equal(t, "numeric_reason_code_enum", newerk.GetReasonCodeFieldName())
}

// TestNewError tests creating errors with enum types and metadata
//
// TestNewError 测试使用枚举类型和 metadata 创建错误
func TestNewError(t *testing.T) {
	t.Run("basic error", func(t *testing.T) {
		erk := newerk.NewError(404, errorspb.ErrorReason_UNKNOWN, "test error")

		require.NotNil(t, erk)
		require.Equal(t, int32(404), erk.Code)
		require.Equal(t, "UNKNOWN", erk.Reason)
		require.Equal(t, "test error", erk.Message)
		require.Equal(t, "0", erk.Metadata["numeric_reason_code_enum"])
	})

	t.Run("database error", func(t *testing.T) {
		erk := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "database failed")

		require.NotNil(t, erk)
		require.Equal(t, int32(500), erk.Code)
		require.Equal(t, "SERVER_DB_ERROR", erk.Reason)
		require.Equal(t, "database failed", erk.Message)
		require.Equal(t, "50001", erk.Metadata["numeric_reason_code_enum"])
	})

	t.Run("with format args", func(t *testing.T) {
		erk := newerk.NewError(404, errorspb.ErrorReason_UNKNOWN, "user %d not found in %s", 123, "database")

		require.NotNil(t, erk)
		require.Equal(t, int32(404), erk.Code)
		require.Equal(t, "user 123 not found in database", erk.Message)
		require.Equal(t, "0", erk.Metadata["numeric_reason_code_enum"])
	})
}

// TestIsError tests checking if error matches enum and HTTP code
//
// TestIsError 测试检查错误是否匹配枚举和 HTTP 代码
func TestIsError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := newerk.IsError(nil, errorspb.ErrorReason_UNKNOWN, 500)
		require.False(t, result)
	})

	t.Run("match success", func(t *testing.T) {
		erk := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "test")
		result := newerk.IsError(erk, errorspb.ErrorReason_SERVER_DB_ERROR, 500)
		require.True(t, result)
	})

	t.Run("wrong reason", func(t *testing.T) {
		erk := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "test")
		result := newerk.IsError(erk, errorspb.ErrorReason_UNKNOWN, 500)
		require.False(t, result)
	})

	t.Run("wrong code", func(t *testing.T) {
		erk := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "test")
		result := newerk.IsError(erk, errorspb.ErrorReason_SERVER_DB_ERROR, 404)
		require.False(t, result)
	})

	t.Run("both wrong", func(t *testing.T) {
		erk := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "test")
		result := newerk.IsError(erk, errorspb.ErrorReason_UNKNOWN, 404)
		require.False(t, result)
	})
}

func TestMetadataEnumNumber(t *testing.T) {
	// Verify different enum values produce correct numeric metadata
	// 验证不同枚举值产生正确的数字 metadata
	erk1 := newerk.NewError(500, errorspb.ErrorReason_UNKNOWN, "test")
	require.Equal(t, "0", erk1.Metadata["numeric_reason_code_enum"])

	erk2 := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_ERROR, "test")
	require.Equal(t, "50001", erk2.Metadata["numeric_reason_code_enum"])

	erk3 := newerk.NewError(500, errorspb.ErrorReason_SERVER_DB_TRANSACTION_ERROR, "test")
	require.Equal(t, "50002", erk3.Metadata["numeric_reason_code_enum"])
}

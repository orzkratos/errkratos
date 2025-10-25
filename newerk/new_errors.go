package newerk

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ProtoErrorEnum constraint for protobuf generated enum types
//
// ProtoErrorEnum protobuf 生成的枚举类型约束
type ProtoErrorEnum interface {
	String() string
	Number() protoreflect.EnumNumber
}

// Config for error generation
//
// Config 错误生成配置
type Config struct {
	ReasonCodeFieldName string // Metadata field name for storing enum numeric value // 存储枚举数值的 metadata 字段名
}

var defaultConfig = &Config{
	ReasonCodeFieldName: "", // default: no metadata
}

// SetReasonCodeFieldName sets the metadata field name for storing enum numeric value
// This field bridges Kratos string reason and frontend numeric enum
//
// SetReasonCodeFieldName 设置用于存储枚举数值的 metadata 字段名
// 此字段用于桥接 Kratos 字符串 reason 和前端数字枚举
func SetReasonCodeFieldName(fieldName string) {
	defaultConfig.ReasonCodeFieldName = fieldName
}

// GetReasonCodeFieldName returns the current reason code field name
//
// GetReasonCodeFieldName 返回当前的原因码字段名
func GetReasonCodeFieldName() string {
	return defaultConfig.ReasonCodeFieldName
}

// IsError checks if the given error matches the specified enum and HTTP code
//
// IsError 检查给定错误是否匹配指定的枚举和 HTTP 代码
func IsError[E ProtoErrorEnum](err error, enum E, httpCode int) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == enum.String() && e.Code == int32(httpCode)
}

// NewError creates a new error with the specified parameters
// If reason code field name is configured, adds enum numeric value to metadata
//
// NewError 创建带指定参数的新错误
// 如果配置了原因码字段名，则将枚举数值添加到 metadata
func NewError[E ProtoErrorEnum](httpCode int, enum E, format string, args ...interface{}) *errors.Error {
	erk := errors.New(httpCode, enum.String(), fmt.Sprintf(format, args...))
	if defaultConfig.ReasonCodeFieldName != "" {
		erk = erk.WithMetadata(map[string]string{
			defaultConfig.ReasonCodeFieldName: fmt.Sprintf("%d", enum.Number()),
		})
	}
	return erk
}

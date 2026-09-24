package errors

// 业务错误码，前后端共用同一套语义
const (
	CodeInvalidPayload = "INVALID_PAYLOAD"
	CodeValidation     = "VALIDATION_ERROR"
	CodeDuplicateSkill = "DUPLICATE_SKILL"
)

// New 构造一个业务错误，错误码与错误消息集中在此管理
func New(code, message string) *BusinessError {
	return &BusinessError{Code: code, Message: message}
}

package errors

import "net/http"

// 技能发布相关的错误码。
const (
	CodeSkillValidation = "SKILL_VALIDATION_FAILED"
	CodeSkillDuplicate  = "SKILL_DUPLICATE_PUBLISH"
)

// SkillValidation 返回资料不完整或非法时的具体失败原因。
func SkillValidation(message string) *BusinessError {
	return New(CodeSkillValidation, message, http.StatusBadRequest)
}

// SkillDuplicate 返回同一技能重复发布的失败原因。
func SkillDuplicate(title string) *BusinessError {
	return New(CodeSkillDuplicate, "你已经在技能墙发布过「"+title+"」，同一技能不能重复发布", http.StatusConflict)
}

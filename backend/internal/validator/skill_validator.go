package validator

import (
	"strings"

	"cyskillswap/internal/constants"
	apperrors "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func countNonEmpty(values []string) int {
	count := 0
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			count++
		}
	}
	return count
}

// ValidatePublishSkill 校验发布请求，资料不完整时返回具体失败原因
func ValidatePublishSkill(req model.PublishSkillRequest) *apperrors.BusinessError {
	if strings.TrimSpace(req.Title) == "" {
		return apperrors.New(apperrors.CodeValidation, "资料不完整：请填写技能名称")
	}
	if !contains(constants.SkillCategories, req.Category) {
		return apperrors.New(apperrors.CodeValidation, "资料不完整：请选择有效的技能类别")
	}
	if req.Level < constants.SkillLevelMin || req.Level > constants.SkillLevelMax {
		return apperrors.New(apperrors.CodeValidation, "资料不完整：熟练度需在 1-100 之间")
	}
	for _, slot := range req.TimeSlots {
		trimmed := strings.TrimSpace(slot)
		if trimmed != "" && !contains(constants.ExchangeTimeSlots, trimmed) {
			return apperrors.New(apperrors.CodeValidation, "可交换时段包含无效选项："+trimmed)
		}
	}
	if countNonEmpty(req.TimeSlots) == 0 {
		return apperrors.New(apperrors.CodeValidation, "资料不完整：请至少选择一个可交换时段")
	}
	if countNonEmpty(req.WantedSkills) == 0 {
		return apperrors.New(apperrors.CodeValidation, "资料不完整：请至少填写一个想学技能")
	}
	return nil
}

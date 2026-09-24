package validator

import (
	"strings"

	"cyskillswap/internal/constants"
	cerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

const (
	minSkillLevel = 1
	maxSkillLevel = 100
)

// ValidatePublishSkill 校验发布表单，资料不完整时返回具体失败原因。
func ValidatePublishSkill(req model.PublishSkillRequest) *cerr.BusinessError {
	if strings.TrimSpace(req.Title) == "" {
		return cerr.SkillValidation("请填写技能名称")
	}
	if strings.TrimSpace(req.Category) == "" {
		return cerr.SkillValidation("请选择技能类别")
	}
	if !contains(constants.SkillCategories, req.Category) {
		return cerr.SkillValidation("技能类别「" + req.Category + "」不在可选范围内")
	}
	if req.Level < minSkillLevel || req.Level > maxSkillLevel {
		return cerr.SkillValidation("请设置 1-100 之间的熟练度")
	}
	if strings.TrimSpace(req.Campus) == "" {
		return cerr.SkillValidation("请选择所在校区")
	}
	if !contains(constants.Campuses, req.Campus) {
		return cerr.SkillValidation("校区「" + req.Campus + "」不在可选范围内")
	}
	if len(req.TimeSlots) == 0 {
		return cerr.SkillValidation("请至少选择一个可交换时段")
	}
	for _, slot := range req.TimeSlots {
		if !contains(constants.ExchangeTimeSlots, slot) {
			return cerr.SkillValidation("可交换时段包含无效选项：" + slot)
		}
	}
	if len(trimmed(req.WantedSkills)) == 0 {
		return cerr.SkillValidation("请至少填写一个想学的技能方向")
	}
	return nil
}

func contains(options []string, value string) bool {
	for _, option := range options {
		if option == value {
			return true
		}
	}
	return false
}

func trimmed(values []string) []string {
	result := []string{}
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			result = append(result, strings.TrimSpace(value))
		}
	}
	return result
}

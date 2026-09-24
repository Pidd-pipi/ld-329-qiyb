package service

import (
	"fmt"
	"strings"

	"cyskillswap/internal/constants"
	apperrors "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
	"cyskillswap/internal/validator"
)

func trimNonEmpty(values []string) []string {
	trimmed := make([]string, 0, len(values))
	for _, value := range values {
		if s := strings.TrimSpace(value); s != "" {
			trimmed = append(trimmed, s)
		}
	}
	return trimmed
}

// PublishSkill 校验并发布当前用户的技能，随后立即生成互补匹配；
// 没有合适人选时保留发布并返回等待提示
func PublishSkill(req model.PublishSkillRequest) (model.PublishSkillResponse, *apperrors.BusinessError) {
	if err := validator.ValidatePublishSkill(req); err != nil {
		logger.Warn("skill publish rejected:", err.Message)
		return model.PublishSkillResponse{}, err
	}
	title := strings.TrimSpace(req.Title)
	if repository.SkillExists(constants.CurrentUser, title) {
		err := apperrors.New(apperrors.CodeDuplicateSkill,
			fmt.Sprintf("你已发布过技能「%s」，同一技能请勿重复发布", title))
		logger.Warn("skill publish rejected:", err.Message)
		return model.PublishSkillResponse{}, err
	}
	description := strings.TrimSpace(req.Description)
	if description == "" {
		description = constants.DefaultSkillDescription
	}
	skill := repository.AddSkill(model.Skill{
		Owner:        constants.CurrentUser,
		Title:        title,
		Category:     req.Category,
		Level:        req.Level,
		Campus:       constants.CurrentUserCampus,
		Description:  description,
		TimeSlots:    trimNonEmpty(req.TimeSlots),
		WantedSkills: trimNonEmpty(req.WantedSkills),
		Rewards:      []string{constants.RewardSkillSwap},
		Portfolio:    constants.DefaultPortfolio,
	})
	matches := repository.AddMatches(findMatchesForSkill(skill, repository.ListSkills(), repository.ListNeedsByRequester))
	waiting := len(matches) == 0
	message := fmt.Sprintf("发布成功，已为你匹配到 %d 位时间有交集且技能互补的同学。", len(matches))
	if waiting {
		message = constants.PublishWaitingMessage
	}
	logger.Info("skill published:", skill.Title, "matches:", len(matches))
	return model.PublishSkillResponse{Skill: skill, Matches: matches, Waiting: waiting, Message: message}, nil
}

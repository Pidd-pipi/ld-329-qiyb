package service

import (
	"fmt"
	"strings"

	"cyskillswap/internal/algorithm"
	"cyskillswap/internal/constants"
	cerr "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
	"cyskillswap/internal/validator"
)

// PublishSkill 校验并发布新技能，随后立即为其生成互补匹配。
func PublishSkill(req model.PublishSkillRequest) (model.PublishSkillResponse, *cerr.BusinessError) {
	req.Owner = strings.TrimSpace(req.Owner)
	if req.Owner == "" {
		req.Owner = constants.CurrentUserName
	}
	req.Title = strings.TrimSpace(req.Title)

	if verr := validator.ValidatePublishSkill(req); verr != nil {
		logger.Warn("publish skill rejected:", verr.Message)
		return model.PublishSkillResponse{}, verr
	}

	if _, duplicated := repository.FindSkillByOwnerAndTitle(req.Owner, req.Title); duplicated {
		derr := cerr.SkillDuplicate(req.Title)
		logger.Warn("publish skill rejected:", derr.Message)
		return model.PublishSkillResponse{}, derr
	}

	skill := repository.AddSkill(model.Skill{
		Owner:        req.Owner,
		Title:        req.Title,
		Category:     req.Category,
		Level:        req.Level,
		Campus:       req.Campus,
		Description:  strings.TrimSpace(req.Description),
		TimeSlots:    req.TimeSlots,
		WantedSkills: req.WantedSkills,
		Rewards:      req.Rewards,
		Portfolio:    strings.TrimSpace(req.Portfolio),
		Status:       constants.SkillStatusWaiting,
	})

	matches := algorithm.FindMatchesForSkill(skill, repository.ListSkills())
	saved := make([]model.Match, 0, len(matches))
	for _, match := range matches {
		saved = append(saved, repository.AddMatch(match))
	}

	response := model.PublishSkillResponse{Skill: skill, Matches: saved}
	if len(saved) == 0 {
		response.Status = constants.PublishStatusWaiting
		response.Message = constants.PublishWaitingMessage
	} else {
		response.Skill = repository.UpdateSkillStatus(skill.ID, constants.SkillStatusMatched)
		response.Status = constants.PublishStatusMatched
		response.Message = fmt.Sprintf("发布成功，为你匹配到 %d 位时间交集且技能互补的同学", len(saved))
	}
	logger.Info("skill published:", skill.Owner, skill.Title, "matches:", len(saved))
	return response, nil
}

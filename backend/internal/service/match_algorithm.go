package service

import (
	"fmt"
	"strings"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// wantedSkillHit 返回候选技能命中的想学技能，未命中返回空串
func wantedSkillHit(wantedSkills []string, candidate model.Skill) string {
	for _, wanted := range wantedSkills {
		if strings.Contains(candidate.Title, wanted) ||
			strings.Contains(candidate.Category, wanted) ||
			strings.Contains(wanted, candidate.Category) {
			return wanted
		}
	}
	return ""
}

// needsSkill 判断对方是否发布过需要该技能类别的需求（互补的另一侧）
func needsSkill(needs []model.Need, category string) bool {
	for _, need := range needs {
		if need.Category == category || strings.Contains(need.Title, category) {
			return true
		}
	}
	return false
}

// commonTimeSlots 计算双方可交换时段的交集
func commonTimeSlots(mine, theirs []string) []string {
	theirsSet := make(map[string]bool, len(theirs))
	for _, slot := range theirs {
		theirsSet[slot] = true
	}
	common := make([]string, 0)
	for _, slot := range mine {
		if theirsSet[slot] {
			common = append(common, slot)
		}
	}
	return common
}

// matchScore 匹配度 = 基础分 + 共同时段加分 + 对方熟练度折算分，封顶 99
func matchScore(level, sharedSlots int) int {
	score := constants.MatchBaseScore + constants.MatchSlotBonus*sharedSlots + level/constants.MatchLevelBonusFactor
	if score > constants.MatchMaxScore {
		return constants.MatchMaxScore
	}
	return score
}

// findMatchesForSkill 只为时间有交集且彼此技能互补的同学生成匹配：
// 对方提供我想学的技能，且对方发布的需求需要我提供的技能类别
func findMatchesForSkill(skill model.Skill, candidates []model.Skill, needsOf func(string) []model.Need) []model.Match {
	matches := make([]model.Match, 0)
	for _, candidate := range candidates {
		if candidate.Owner == skill.Owner {
			continue
		}
		shared := commonTimeSlots(skill.TimeSlots, candidate.TimeSlots)
		if len(shared) == 0 {
			continue
		}
		hit := wantedSkillHit(skill.WantedSkills, candidate)
		if hit == "" {
			continue
		}
		if !needsSkill(needsOf(candidate.Owner), skill.Category) {
			continue
		}
		matches = append(matches, model.Match{
			Provider:    candidate.Owner,
			Learner:     skill.Owner,
			OfferSkill:  candidate.Title,
			WantedSkill: skill.Title,
			Score:       matchScore(candidate.Level, len(shared)),
			CommonSlots: shared,
			Recommendation: fmt.Sprintf("你想学「%s」，%s 的「%s」正好互补；对方也需要你的「%s」，共同时段已对齐。",
				hit, candidate.Owner, candidate.Title, skill.Category),
		})
	}
	return matches
}

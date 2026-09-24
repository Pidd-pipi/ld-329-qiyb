package algorithm

import (
	"strings"

	"cyskillswap/internal/model"
)

// FindMatchesForSkill 只为时间有交集且彼此技能互补的同学生成匹配。
func FindMatchesForSkill(skill model.Skill, candidates []model.Skill) []model.Match {
	matches := []model.Match{}
	for _, candidate := range candidates {
		if candidate.ID == skill.ID || candidate.Owner == skill.Owner {
			continue
		}
		common := CommonSlots(skill.TimeSlots, candidate.TimeSlots)
		if len(common) == 0 {
			continue
		}
		if !IsComplementary(skill, candidate) {
			continue
		}
		matches = append(matches, model.Match{
			Provider:       candidate.Owner,
			Learner:        skill.Owner,
			OfferSkill:     candidate.Title,
			WantedSkill:    skill.Title,
			Score:          CalculateScore(skill, candidate, common),
			CommonSlots:    common,
			Recommendation: BuildRecommendation(skill, candidate, common),
		})
	}
	return matches
}

// CommonSlots 计算两名同学可交换时段的交集。
func CommonSlots(a, b []string) []string {
	inB := map[string]bool{}
	for _, slot := range b {
		inB[slot] = true
	}
	common := []string{}
	for _, slot := range a {
		if inB[slot] {
			common = append(common, slot)
		}
	}
	return common
}

// IsComplementary 判断双方是否彼此技能互补：
// 对方提供的正是我想学的，且我提供的也正是对方想学的。
func IsComplementary(a, b model.Skill) bool {
	return wantsOffer(a.WantedSkills, b) && wantsOffer(b.WantedSkills, a)
}

// wantsOffer 判断 offered 技能（类别或名称）是否命中 wanted 想学方向。
func wantsOffer(wanted []string, offered model.Skill) bool {
	for _, want := range wanted {
		want = strings.TrimSpace(want)
		if want == "" {
			continue
		}
		if strings.EqualFold(want, offered.Category) ||
			strings.Contains(offered.Title, want) ||
			strings.Contains(want, offered.Category) {
			return true
		}
	}
	return false
}

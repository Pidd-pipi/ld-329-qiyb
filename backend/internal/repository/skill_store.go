package repository

import (
	"strings"

	"cyskillswap/internal/model"
)

func normalizeTitle(title string) string {
	return strings.ToLower(strings.TrimSpace(title))
}

// SkillExists 判断该用户是否已发布过同名技能，用于拦截重复发布
func SkillExists(owner, title string) bool {
	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, skill := range store.skills {
		if skill.Owner == owner && normalizeTitle(skill.Title) == normalizeTitle(title) {
			return true
		}
	}
	return false
}

// AddSkill 写入新技能并分配自增 ID，返回落库后的技能
func AddSkill(skill model.Skill) model.Skill {
	store.mu.Lock()
	defer store.mu.Unlock()
	skill.ID = store.nextSkillID
	store.nextSkillID++
	store.skills = append(store.skills, skill)
	return skill
}

// ListSkillsByOwner 返回某位用户发布的全部技能（技能墙数据来源）
func ListSkillsByOwner(owner string) []model.Skill {
	store.mu.RLock()
	defer store.mu.RUnlock()
	owned := make([]model.Skill, 0)
	for _, skill := range store.skills {
		if skill.Owner == owner {
			owned = append(owned, skill)
		}
	}
	return owned
}

// ListNeedsByRequester 返回某位用户发布过的全部需求（判断互补时使用）
func ListNeedsByRequester(requester string) []model.Need {
	store.mu.RLock()
	defer store.mu.RUnlock()
	needs := make([]model.Need, 0)
	for _, need := range store.needs {
		if need.Requester == requester {
			needs = append(needs, need)
		}
	}
	return needs
}

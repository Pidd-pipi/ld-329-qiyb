package algorithm

import (
	"strings"

	"cyskillswap/internal/model"
)

// 匹配度评分的统一规则：基础分 + 共同时段加分 + 熟练度接近加分。
const (
	baseScore        = 70
	slotBonusEach    = 8
	slotBonusMax     = 16
	levelBonusMax    = 14
	scoreCap         = 99
	levelScaleFactor = 100
)

// CalculateScore 根据共同时段数量和双方熟练度差距计算匹配度。
func CalculateScore(a, b model.Skill, common []string) int {
	score := baseScore

	slotBonus := len(common) * slotBonusEach
	if slotBonus > slotBonusMax {
		slotBonus = slotBonusMax
	}
	score += slotBonus

	diff := a.Level - b.Level
	if diff < 0 {
		diff = -diff
	}
	levelBonus := levelBonusMax - diff*levelBonusMax/levelScaleFactor
	if levelBonus < 0 {
		levelBonus = 0
	}
	score += levelBonus

	if score > scoreCap {
		score = scoreCap
	}
	return score
}

// BuildRecommendation 生成匹配推荐理由，说明互补关系和共同时段。
func BuildRecommendation(a, b model.Skill, common []string) string {
	return b.Owner + " 提供的「" + b.Title + "」正是你想学的方向，" +
		"你的「" + a.Title + "」也正是对方想学的；" +
		"共同可交换时段：" + strings.Join(common, "、") + "。"
}

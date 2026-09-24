package repository

import (
	"sync"

	"cyskillswap/internal/model"
)

var (
	matchMu     sync.RWMutex
	matchStore  []model.Match
	nextMatchID int
)

func init() {
	matchStore = []model.Match{
		{ID: 1, Provider: "林澈", Learner: "孟野", OfferSkill: "毕业照人像摄影", WantedSkill: "民谣吉他陪练", Score: 96, CommonSlots: []string{"周三晚", "周六上午"}, Recommendation: "互补技能明确，双方均接受技能交换。"},
		{ID: 2, Provider: "周芮", Learner: "许安", OfferSkill: "Python 数据分析", WantedSkill: "论文数据清洗", Score: 89, CommonSlots: []string{"周二晚"}, Recommendation: "时间匹配且需求描述命中 pandas/可视化。"},
		{ID: 3, Provider: "孟野", Learner: "林澈", OfferSkill: "民谣吉他陪练", WantedSkill: "宣传照拍摄", Score: 91, CommonSlots: []string{"周六上午"}, Recommendation: "互换回报类型一致，信用分权重较高。"},
	}
	nextMatchID = len(matchStore) + 1
}

// ListMatches 返回系统已生成的全部匹配。
func ListMatches() []model.Match {
	matchMu.RLock()
	defer matchMu.RUnlock()
	result := make([]model.Match, len(matchStore))
	copy(result, matchStore)
	return result
}

// AddMatch 写入一条新生成的匹配并分配 ID。
func AddMatch(match model.Match) model.Match {
	matchMu.Lock()
	defer matchMu.Unlock()
	match.ID = nextMatchID
	nextMatchID++
	matchStore = append(matchStore, match)
	return match
}

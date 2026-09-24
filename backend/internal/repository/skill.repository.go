package repository

import (
	"strings"
	"sync"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

var (
	skillMu     sync.RWMutex
	skillStore  []model.Skill
	nextSkillID int
)

func init() {
	skillStore = []model.Skill{
		{ID: 1, Owner: "林澈", Title: "毕业照人像摄影", Category: "摄影", Level: 92, Campus: "东校区", Description: "提供构图、修图和毕业季跟拍，可交换吉他入门课。", TimeSlots: []string{"周三晚", "周六上午"}, WantedSkills: []string{"乐器"}, Rewards: []string{"技能交换", "请吃饭"}, Portfolio: "12组校园人像作品", Status: constants.SkillStatusMatched},
		{ID: 2, Owner: "周芮", Title: "Python 数据分析", Category: "编程", Level: 88, Campus: "中心校区", Description: "pandas、可视化、论文数据清洗辅导，接受小额报酬。", TimeSlots: []string{"周二晚", "周日全天"}, WantedSkills: []string{"摄影"}, Rewards: []string{"技能交换", "小额报酬"}, Portfolio: "3份课程项目证书", Status: constants.SkillStatusMatched},
		{ID: 3, Owner: "孟野", Title: "民谣吉他陪练", Category: "乐器", Level: 81, Campus: "西校区", Description: "节奏型、弹唱和舞台经验分享，想找人拍宣传照。", TimeSlots: []string{"周三晚", "周六上午"}, WantedSkills: []string{"摄影"}, Rewards: []string{"技能交换", "无偿"}, Portfolio: "校园音乐节演出视频", Status: constants.SkillStatusMatched},
	}
	nextSkillID = len(skillStore) + 1
}

// ListSkills 返回技能墙上的全部技能。
func ListSkills() []model.Skill {
	skillMu.RLock()
	defer skillMu.RUnlock()
	result := make([]model.Skill, len(skillStore))
	copy(result, skillStore)
	return result
}

// ListSkillsByOwner 返回某位同学发布的技能（用于个人技能墙）。
func ListSkillsByOwner(owner string) []model.Skill {
	skillMu.RLock()
	defer skillMu.RUnlock()
	result := []model.Skill{}
	for _, skill := range skillStore {
		if skill.Owner == owner {
			result = append(result, skill)
		}
	}
	return result
}

// FindSkillByOwnerAndTitle 判断同一同学是否已发布过同名技能。
func FindSkillByOwnerAndTitle(owner, title string) (model.Skill, bool) {
	skillMu.RLock()
	defer skillMu.RUnlock()
	normalized := strings.TrimSpace(title)
	for _, skill := range skillStore {
		if skill.Owner == owner && strings.EqualFold(strings.TrimSpace(skill.Title), normalized) {
			return skill, true
		}
	}
	return model.Skill{}, false
}

// AddSkill 把新技能写入技能墙并分配 ID。
func AddSkill(skill model.Skill) model.Skill {
	skillMu.Lock()
	defer skillMu.Unlock()
	skill.ID = nextSkillID
	nextSkillID++
	skillStore = append(skillStore, skill)
	return skill
}

// UpdateSkillStatus 更新技能墙上某个技能的匹配状态，返回更新后的技能。
func UpdateSkillStatus(id int, status string) model.Skill {
	skillMu.Lock()
	defer skillMu.Unlock()
	for i := range skillStore {
		if skillStore[i].ID == id {
			skillStore[i].Status = status
			return skillStore[i]
		}
	}
	return model.Skill{}
}

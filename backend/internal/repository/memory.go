package repository

import (
	"sync"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// memoryStore 以内存方式保存演示数据，发布后立即对后续读取可见
type memoryStore struct {
	mu           sync.RWMutex
	skills       []model.Skill
	needs        []model.Need
	matches      []model.Match
	appointments []model.Appointment
	reviews      []model.Review
	messages     []model.Conversation
	nextSkillID  int
	nextMatchID  int
}

var store = newMemoryStore()

func newMemoryStore() *memoryStore {
	return &memoryStore{
		skills: []model.Skill{
			{ID: 1, Owner: "林澈", Title: "毕业照人像摄影", Category: "摄影", Level: 92, Campus: "东校区", Description: "提供构图、修图和毕业季跟拍，可交换吉他入门课。", TimeSlots: []string{"周三晚", "周六上午"}, WantedSkills: []string{"吉他"}, Rewards: []string{"技能交换", "请吃饭"}, Portfolio: "12组校园人像作品"},
			{ID: 2, Owner: "周芮", Title: "Python 数据分析", Category: "编程", Level: 88, Campus: "中心校区", Description: "pandas、可视化、论文数据清洗辅导，接受小额报酬。", TimeSlots: []string{"周二晚", "周日全天"}, WantedSkills: []string{"平面设计"}, Rewards: []string{"技能交换", "小额报酬"}, Portfolio: "3份课程项目证书"},
			{ID: 3, Owner: "孟野", Title: "民谣吉他陪练", Category: "乐器", Level: 81, Campus: "西校区", Description: "节奏型、弹唱和舞台经验分享，想找人拍宣传照。", TimeSlots: []string{"周三晚", "周六上午"}, WantedSkills: []string{"摄影"}, Rewards: []string{"技能交换", "无偿"}, Portfolio: "校园音乐节演出视频"},
		},
		needs: []model.Need{
			{ID: 1, Requester: "孟野", Title: "找人帮忙拍乐队宣传照", Category: "摄影", Campus: "西校区", ExpectTime: "本周六上午", BudgetType: "技能交换", Description: "可交换 3 次吉他课，希望会调色和室外构图。", Responses: 5},
			{ID: 2, Requester: "许安", Title: "求教 Python 数据分析", Category: "编程", Campus: "中心校区", ExpectTime: "周二晚", BudgetType: "小额报酬", Description: "论文问卷数据需要清洗和画图，最好有 pandas 经验。", Responses: 8},
			{ID: 3, Requester: "林澈", Title: "想学吉他扫弦入门", Category: "乐器", Campus: "东校区", ExpectTime: "周三晚", BudgetType: "技能交换", Description: "用摄影课交换吉他基础，希望同校区或线上。", Responses: 3},
		},
		matches: []model.Match{
			{ID: 1, Provider: "林澈", Learner: "孟野", OfferSkill: "毕业照人像摄影", WantedSkill: "民谣吉他陪练", Score: 96, CommonSlots: []string{"周三晚", "周六上午"}, Recommendation: "互补技能明确，双方均接受技能交换。"},
			{ID: 2, Provider: "周芮", Learner: "许安", OfferSkill: "Python 数据分析", WantedSkill: "论文数据清洗", Score: 89, CommonSlots: []string{"周二晚"}, Recommendation: "时间匹配且需求描述命中 pandas/可视化。"},
			{ID: 3, Provider: "孟野", Learner: "林澈", OfferSkill: "民谣吉他陪练", WantedSkill: "宣传照拍摄", Score: 91, CommonSlots: []string{"周六上午"}, Recommendation: "互换回报类型一致，信用分权重较高。"},
		},
		appointments: []model.Appointment{
			{ID: 1, Pair: "林澈 ↔ 孟野", Time: "周六 10:00", Place: "东校区湖边", Status: "双方已确认", Agenda: "先拍宣传照，再约 2 次吉他课"},
			{ID: 2, Pair: "周芮 ↔ 许安", Time: "周二 19:30", Place: "线上会议室", Status: "等待对方确认", Agenda: "导入问卷 CSV 并完成基础可视化"},
		},
		reviews: []model.Review{
			{ID: 1, From: "孟野", To: "林澈", Rating: 5, Content: "构图建议很细，成片当天就给了预览。"},
			{ID: 2, From: "林澈", To: "孟野", Rating: 5, Content: "吉他入门节奏拆得很清楚，课后还发了练习谱。"},
		},
		messages: []model.Conversation{
			{ID: 1, WithUser: "孟野", Unread: 2, Messages: []string{"周六湖边光线不错", "我带两套衣服可以吗？"}},
			{ID: 2, WithUser: "系统通知", Unread: 1, Messages: []string{"你与周芮的 Python 数据分析预约待确认。"}},
		},
		nextSkillID: 4,
		nextMatchID: 4,
	}
}

func ListSkills() []model.Skill {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Skill(nil), store.skills...)
}

func ListNeeds() []model.Need {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Need(nil), store.needs...)
}

func ListMatches() []model.Match {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Match(nil), store.matches...)
}

func ListAppointments() []model.Appointment {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Appointment(nil), store.appointments...)
}

func ListReviews() []model.Review {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Review(nil), store.reviews...)
}

func ListMessages() []model.Conversation {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return append([]model.Conversation(nil), store.messages...)
}

// GetProfile 个人主页：技能墙取当前用户已发布的技能，雷达图合并各项技能熟练度
func GetProfile() model.Profile {
	wall := ListSkillsByOwner(constants.CurrentUser)
	radar := map[string]int{"摄影": 92, "修图": 86, "沟通": 90, "编程": 42, "乐器": 35}
	for _, skill := range wall {
		if skill.Level > radar[skill.Category] {
			radar[skill.Category] = skill.Level
		}
	}
	return model.Profile{
		Name: constants.CurrentUser, Major: "新闻传播 2023", CreditScore: 91, CreditLevel: constants.CreditGold,
		SkillWall: wall,
		Radar:     radar,
		History:   []string{"完成毕业照拍摄交换", "响应 Python 数据分析需求", "预约吉他入门课"},
		Reviews:   ListReviews(),
	}
}

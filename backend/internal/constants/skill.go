package constants

// 当前登录用户（JWT 接入前的默认身份）及其校区
const (
	CurrentUser       = "林澈"
	CurrentUserCampus = "东校区"
)

// ExchangeTimeSlots 可交换时段选项
var ExchangeTimeSlots = []string{"周一晚", "周二晚", "周三晚", "周四晚", "周五晚", "周六上午", "周六下午", "周日上午", "周日下午", "周日全天"}

// 熟练度取值范围
const (
	SkillLevelMin = 1
	SkillLevelMax = 100
)

// 匹配度计算规则：基础分 + 每个共同时段加分 + 对方熟练度折算分，封顶 99
const (
	MatchBaseScore        = 70
	MatchSlotBonus        = 6
	MatchLevelBonusFactor = 10
	MatchMaxScore         = 99
)

// 发布技能时的默认填充值
const (
	RewardSkillSwap         = "技能交换"
	DefaultPortfolio        = "暂未上传作品"
	DefaultSkillDescription = "暂无补充描述"
)

// PublishWaitingMessage 没有合适人选时的等待提示
const PublishWaitingMessage = "发布成功，暂无时间有交集且技能互补的同学，已为你保留发布，仍在等待匹配。"

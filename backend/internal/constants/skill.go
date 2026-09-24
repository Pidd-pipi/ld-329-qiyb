package constants

// CurrentUserName 是当前登录学生（演示环境固定为林澈）。
const CurrentUserName = "林澈"

// ExchangeTimeSlots 是平台统一维护的可交换时段选项。
var ExchangeTimeSlots = []string{
	"周一晚", "周二晚", "周三晚", "周四晚", "周五晚",
	"周六上午", "周六下午", "周日上午", "周日下午", "周日全天",
}

// Campuses 是可选校区列表。
var Campuses = []string{"东校区", "西校区", "中心校区"}

// RewardTypes 是期望回报类型选项。
var RewardTypes = []string{"技能交换", "小额报酬", "请吃饭", "无偿"}

// 技能在技能墙上的匹配状态。
const (
	SkillStatusMatched = "已匹配"
	SkillStatusWaiting = "等待匹配"
)

// 发布结果状态。
const (
	PublishStatusMatched = "matched"
	PublishStatusWaiting = "waiting"
)

// PublishWaitingMessage 是没有合适人选时给发布者的说明。
const PublishWaitingMessage = "技能已发布到技能墙，暂时没有时间和技能都互补的同学，仍在等待合适人选"

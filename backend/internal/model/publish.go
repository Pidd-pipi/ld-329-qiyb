package model

// PublishSkillRequest 是发布技能时前端提交的表单内容。
type PublishSkillRequest struct {
	Owner        string   `json:"owner"`
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Level        int      `json:"level"`
	Campus       string   `json:"campus"`
	Description  string   `json:"description"`
	TimeSlots    []string `json:"timeSlots"`
	WantedSkills []string `json:"wantedSkills"`
	Rewards      []string `json:"rewards"`
	Portfolio    string   `json:"portfolio"`
}

// PublishSkillResponse 返回新发布的技能、即时匹配结果和发布状态。
type PublishSkillResponse struct {
	Skill   Skill   `json:"skill"`
	Matches []Match `json:"matches"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}

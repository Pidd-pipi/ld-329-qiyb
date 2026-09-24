package model

// PublishSkillRequest 发布技能的请求体：技能名称、熟练度、可交换时段和想学技能为必填
type PublishSkillRequest struct {
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Level        int      `json:"level"`
	TimeSlots    []string `json:"timeSlots"`
	WantedSkills []string `json:"wantedSkills"`
	Description  string   `json:"description"`
}

// PublishSkillResponse 发布结果：新技能、命中的匹配以及是否仍在等待
type PublishSkillResponse struct {
	Skill   Skill   `json:"skill"`
	Matches []Match `json:"matches"`
	Waiting bool    `json:"waiting"`
	Message string  `json:"message"`
}

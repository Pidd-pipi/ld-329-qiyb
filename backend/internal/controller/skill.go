package controller

import (
	"net/http"

	apperrors "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// PublishSkill 发布当前用户的技能：成功返回新技能与匹配结果，失败返回具体原因
func PublishSkill(c *gin.Context) {
	var req model.PublishSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("skill publish bind failed:", err.Error())
		c.JSON(http.StatusBadRequest, apperrors.New(apperrors.CodeInvalidPayload, "提交内容格式不正确，请检查后重试"))
		return
	}
	result, berr := service.PublishSkill(req)
	if berr != nil {
		c.JSON(http.StatusBadRequest, berr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

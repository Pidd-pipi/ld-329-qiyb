package controller

import (
	"net/http"

	cerr "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// PublishSkill 处理技能发布请求，成功返回新技能和即时匹配结果。
func PublishSkill(c *gin.Context) {
	var req model.PublishSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, cerr.SkillValidation("提交内容格式不正确，请检查表单后重试"))
		return
	}
	response, berr := service.PublishSkill(req)
	if berr != nil {
		c.JSON(berr.HTTPStatus, berr)
		return
	}
	c.JSON(http.StatusCreated, response)
}

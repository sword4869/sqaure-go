package controller

import (
	"net/http"
	"test/store"

	"github.com/gin-gonic/gin"
)

// ListTags 获取标签列表
// @Router /tags/list [post]
// @Summary 获取标签列表
// @Description 获取标签列表
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
func ListTags(ctx *gin.Context) {
	// 获取标签列表
	tags, err := store.NewTag().ListAllTags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取标签列表失败: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"list": tags,
	})
}

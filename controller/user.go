package controller

import (
	"net/http"
	"test/store"

	"github.com/gin-gonic/gin"
)

// @Router /user/new_random_user [post]
// @Summary 获取随机新用户
// @Description 获取随机新用户
// @Tags         user
// @Accept json
// @Produce json
// @Success 200 {object} store.UserRsp
func NewRandomUser(c *gin.Context) {
	user := store.NewUser()
	err := user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userRsp, err := user.GetUserAvatarById(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRsp)
}

type GetUserByIDRequest struct {
	Id int64 `json:"id" binding:"required" example:"1"` // 用户ID
}

// @Router /user/get_by_id [post]
// @Summary 获取指定用户
// @Description 获取指定用户
// @Tags         user
// @Accept json
// @Produce json
// @Param        id  body  GetUserByIDRequest  true  "获取指定用户"
// @Success 200 {object} store.UserRsp
func GetUserByID(c *gin.Context) {
	var req GetUserByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRsp, err := store.NewUser().GetUserAvatarById(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRsp)
}

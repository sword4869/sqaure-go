package controller

import (
	"net/http"
	"test/store"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostLikeGetParams struct {
	// 帖子ID
	PostId int `json:"post_id" form:"post_id" binding:"required" example:"1"`
	// 用户ID
	UserId int `json:"user_id" form:"user_id" binding:"required" example:"1"`
}

type PostLikeGetResponse struct {
	// 是否点赞: 0 未点赞, 1 已点赞
	IsLike int `json:"is_like" example:"1"`
}

// @Summary 获取当前用户对指定帖子的点赞状态
// @Description 0 未点赞, 1 已点赞
// @Router /post_like/get [post]
// @Tags post_like
// @Param p body PostLikeGetParams true "参数"
// @Success 200 {object} PostLikeGetResponse "返回结果"
func PostLikeGet(ctx *gin.Context) {
	var params PostLikeGetParams
	if err := ctx.ShouldBind(&params); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}

	postLike, err := store.NewPostLike().GetByUserIdAndPostId(params.UserId, params.PostId)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取发帖点赞失败: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, PostLikeGetResponse{IsLike: postLike.IsLike})
}

type PostLikeUpdateLikedParams struct {
	PostId int `json:"post_id" form:"post_id" binding:"required" example:"1"` // 帖子ID
	UserId int `json:"user_id" form:"user_id" binding:"required" example:"1"` // 用户ID
	IsLike int `json:"is_like" form:"is_like" example:"1"`                    // 是否点赞: 0 取消点赞, 1 点赞
}

// @Summary 更新当前用户对指定帖子的点赞状态
// @Description 0 取消点赞, 1 点赞
// @Router /post_like/update_liked [post]
// @Tags post_like
// @Accept json
// @Produce json
// @Param p body PostLikeUpdateLikedParams true "参数"
func PostLikeUpdateLiked(ctx *gin.Context) {
	var params PostLikeUpdateLikedParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}

	// 点赞总数
	if params.IsLike == 1 {
		err := store.NewPost().Like(params.PostId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "点赞失败: " + err.Error(),
			})
			return
		}
	} else {
		err := store.NewPost().UnLike(params.PostId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "取消点赞失败: " + err.Error(),
			})
			return
		}
	}

	// 用户点赞状态
	err := store.NewPostLike().UpdateLiked(params.UserId, params.PostId, params.IsLike)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取发帖点赞失败: " + err.Error(),
		})
		return
	}

	// 新的点赞总数
	post, err := store.NewPost().GetById(params.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取发帖失败: " + err.Error(),
		})
		return
	}
	if params.IsLike == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "取消点赞成功", "like_count": post.LikeCount})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "点赞成功", "like_count": post.LikeCount})
	}
}

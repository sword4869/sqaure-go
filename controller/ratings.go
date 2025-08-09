package controller

import (
	"net/http"
	"test/store"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateRatingsParams struct {
	UserId int `json:"user_id" form:"user_id" binding:"required" example:"1"` // 用户ID
	PostId int `json:"post_id" form:"post_id" binding:"required" example:"1"` // 帖子ID
	Rating int `json:"rating" form:"rating" binding:"required" example:"5"`   // 评分 -5 ~ 5
}

// @Router /ratings/create [post]
// @Summary 用户对帖子评分
// @Description 用户对帖子评分
// @Tags ratings
// @Accept json
// @Produce json
// @Param p body CreateRatingsParams true "参数示例"
func CreateRatings(ctx *gin.Context) {
	logger := logrus.WithField("func", "CreateRatings")
	var params CreateRatingsParams
	if err := ctx.ShouldBind(&params); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查评分是否在 -5 ~ 5 之间
	if params.Rating < -5 || params.Rating > 5 {
		logger.Error("评分必须在 -5 ~ 5 之间")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "评分必须在 -5 ~ 5 之间"})
		return
	}

	// 检查用户是否已经评分
	ratings, err := store.NewRatings().GetByUserIdAndPostId(params.UserId, params.PostId)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ratings != nil && ratings.Id != 0 {
		logger.Error("用户已经评分")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户已经评分"})
		return
	}

	// 创建用户对帖子的评分
	ratings = &store.Ratings{
		UserId: params.UserId,
		PostId: params.PostId,
		Rating: params.Rating,
	}
	if err := ratings.Create(); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新帖子评分
	// 获取帖子评分平均值, 如果获取失败, 则创建一个
	ratingsAve, err := store.NewRatingsAve().GetByPostId(params.PostId)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ratingsAve == nil || ratingsAve.PostId == 0 {
		ratingsAve = store.NewRatingsAveByPostId(params.PostId)
	}
	logger.Infof("帖子评分平均值: %+v", ratingsAve)
	ratingsAve.RatingAve = (ratingsAve.RatingAve*float64(ratingsAve.RatingCount) + float64(params.Rating)) / float64(ratingsAve.RatingCount+1)
	ratingsAve.RatingCount = ratingsAve.RatingCount + 1
	logger.Infof("更新后帖子评分平均值: %+v", ratingsAve)
	if err := ratingsAve.Upsert(); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "评分成功"})
}

type GetRatingsParams struct {
	UserId int `json:"user_id" form:"user_id" binding:"required" example:"1"` // 用户ID
	PostId int `json:"post_id" form:"post_id" binding:"required" example:"1"` // 帖子ID
}

type GetRatingsResponse struct {
	UserId int `json:"user_id" form:"user_id" example:"1"` // 用户ID
	PostId int `json:"post_id" form:"post_id" example:"1"` // 帖子ID
	Rating int `json:"rating" form:"rating" example:"5"`   // 评分
}

// @Router /ratings/get [post]
// @Summary 获取用户对帖子的评分
// @Description 获取用户对帖子的评分
// @Tags ratings
// @Accept json
// @Produce json
// @Param p body GetRatingsParams true "参数示例"
// @Success 200 {object} GetRatingsResponse
func GetRatings(ctx *gin.Context) {
	var params GetRatingsParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ratings, err := store.NewRatings().GetByUserIdAndPostId(params.UserId, params.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, GetRatingsResponse{
		UserId: ratings.UserId,
		PostId: ratings.PostId,
		Rating: ratings.Rating,
	})
}

type GetRatingsAveParams struct {
	PostId int `json:"post_id" form:"post_id" binding:"required" example:"1"` // 帖子ID
}

type GetRatingsAveResponse struct {
	PostId      int     `json:"post_id" form:"post_id" example:"1"`           // 帖子ID
	RatingAve   float64 `json:"rating_ave" form:"rating_ave" example:"5.0"`   // 平均评分
	RatingCount int     `json:"rating_count" form:"rating_count" example:"1"` // 评分总人数
}

// @Router /ratings/get_ave [post]
// @Summary 获取帖子评分平均值
// @Description 获取帖子评分平均值
// @Tags ratings
// @Accept json
// @Produce json
// @Param p body GetRatingsAveParams true "参数示例"
// @Success 200 {object} GetRatingsAveResponse
func GetRatingsAve(ctx *gin.Context) {
	var params GetRatingsAveParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ratingsAve, err := store.NewRatingsAve().GetByPostId(params.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, GetRatingsAveResponse{
		PostId:      ratingsAve.PostId,
		RatingAve:   ratingsAve.RatingAve,
		RatingCount: ratingsAve.RatingCount,
	})
}

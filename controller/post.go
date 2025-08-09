package controller

import (
	"net/http"
	"test/store"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePost 创建发帖
// @Summary      创建发帖
// @Description  创建发帖
// @Router       /posts/create [post]
// @Param        post  body  store.Post  true  "发帖信息"
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {object}  store.Post
func CreatePost(ctx *gin.Context) {
	var post store.Post

	if err := ctx.ShouldBind(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}

	if err := post.Create(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建发帖失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "发帖成功",
		"data":    post,
	})
}

type ListPostsParams struct {
	Cursor int64 `json:"cursor" form:"cursor" example:"0"`
}

// ListPosts 获取所有发帖列表
// @Summary      获取所有发帖列表
// @Description  获取所有发帖列表
// @Router       /posts/list [post]
// @Tags         posts
// @Param        post  body  ListPostsParams  true  "发帖信息"
// @Accept       json
// @Produce      json
func ListPosts(ctx *gin.Context) {
	var params ListPostsParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}

	if params.Cursor == 0 {
		params.Cursor = time.Now().Unix()
	}

	posts, err := store.NewPost().ListPosts(params.Cursor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取发帖列表失败: " + err.Error(),
		})
		return
	}

	cursor := int64(0)
	if len(posts) > 0 {
		cursor = posts[len(posts)-1].CreatedAt
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取发帖列表成功",
		"data":    posts,
		"cursor":  cursor,
	})
}

type ListPostsByKeywordParams struct {
	Cursor  int64  `json:"cursor" form:"cursor" example:"0"`
	Keyword string `json:"keyword" form:"keyword" example:""` // 关键字，可选
}

// ListPostsByKeyword 通过关键字获取发帖列表
// @Summary      通过关键字获取发帖列表
// @Description  通过关键字获取发帖列表
// @Router       /posts/list_by_keyword [post]
// @Tags         posts
// @Param        post  body  ListPostsByKeywordParams  true  "通过关键字获取发帖列表信息"
// @Accept       json
// @Produce      json
func ListPostsByKeyword(ctx *gin.Context) {
	var params ListPostsByKeywordParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}

	if params.Cursor == 0 {
		params.Cursor = time.Now().Unix()
	}

	posts, err := store.NewPost().ListPostsByKeyword(params.Keyword, params.Cursor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取用户发帖列表失败: " + err.Error(),
		})
		return
	}

	cursor := int64(0)
	if len(posts) > 0 {
		cursor = posts[len(posts)-1].CreatedAt
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取用户发帖列表成功",
		"data":    posts,
		"cursor":  cursor,
	})
}

type GetPostByIdParams struct {
	PostId int `json:"post_id" form:"post_id" example:"1"` // 发帖ID
}

type GetPostByIdResponse struct {
	Post *store.Post `json:"post" form:"post"` // 发帖
}

// @Summary      获取单个发帖
// @Description  获取单个发帖
// @Router       /posts/get_by_id [post]
// @Tags         posts
// @Param        post  body  GetPostByIdParams  true  "获取单个发帖信息"
// @Success      200  {object}  GetPostByIdResponse
func GetPostById(ctx *gin.Context) {
	var params GetPostByIdParams
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON格式错误: " + err.Error(),
		})
		return
	}
	post, err := store.NewPost().GetById(params.PostId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取发帖失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, GetPostByIdResponse{Post: post})
}

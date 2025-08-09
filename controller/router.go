package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router 设置路由
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// 标签相关路由
	api := router.Group("/api")
	{
		// 用户相关路由
		user := api.Group("/user")
		{
			user.POST("/get_by_id", GetUserByID)
			user.POST("/new_random_user", NewRandomUser)
		}

		// 标签管理
		tags := api.Group("/tags")
		{
			tags.POST("/list", ListTags) // 获取标签列表
		}

		// 发帖相关路由
		posts := api.Group("/posts")
		{
			posts.POST("/create", CreatePost)                  // 创建发帖
			posts.POST("/list", ListPosts)                     // 获取发帖列表
			posts.POST("/list_by_keyword", ListPostsByKeyword) // 通过关键字获取发帖列表
			posts.POST("/get_by_id", GetPostById)              // 获取单个发帖
		}

		// 发帖点赞相关路由
		postLike := api.Group("/post_like")
		{
			postLike.POST("/get", PostLikeGet)                  // 获取点赞状态
			postLike.POST("/update_liked", PostLikeUpdateLiked) // 更新点赞状态
		}

		// 评分相关路由
		ratings := api.Group("/ratings")
		{
			ratings.POST("/create", CreateRatings)  // 用户对帖子评分
			ratings.POST("/get", GetRatings)        // 获取用户对帖子的评分
			ratings.POST("/get_ave", GetRatingsAve) // 获取帖子评分平均值
		}
	}

	return router
}

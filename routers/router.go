package routers

import (
	"github.com/Hallelujah1025/Stroke-Survivors/middleware/jwt"
	"github.com/Hallelujah1025/Stroke-Survivors/routers/api"
	v1 "github.com/Hallelujah1025/Stroke-Survivors/routers/api/v1"
	"github.com/gin-gonic/gin"

	"github.com/Hallelujah1025/Stroke-Survivors/pkg/setting"
)

//InitRouter 初始化路径
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//注册
	r.GET("/register", api.Register)
	//登陆
	r.GET("/auth", api.GetAuth)
	//检查答案
	r.GET("/checkAnswer", api.CheckAnswer)
	// 修改密码
	r.GET("/updateAuth", api.UpdateAuth)
	//反馈
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取某人文章列表
		apiv1.GET("/who/:id", v1.GetArticlesFromSomeone)
		//新建文章
		apiv1.GET("/article", v1.AddArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)

		//获取评论列表
		apiv1.GET("/comments", v1.GetComments)
		//获取指定评论
		apiv1.GET("/comments/:id", v1.GetComment)
		//新建评论
		apiv1.GET("/comment", v1.AddComment)

		//添加到收藏
		apiv1.GET("/favorite", v1.AddFavorite)
		//删除收藏
		apiv1.GET("/favorited", v1.DeleteFavorite)
		//添加关注
		apiv1.GET("/follow", v1.AddFollow)
		//取消关注
		apiv1.GET("/followed", v1.DeleteFollow)
		//病人获取关注的医生列表
		apiv1.GET("/doctor", v1.GetDoctor)
		//医生获取关注自己的病人列表
		apiv1.GET("/patient", v1.GetPatient)
		//获取收藏的文章
		apiv1.GET("/favoriteArticle", v1.GetFavoriteArticle)
		//搜索医生
		apiv1.GET("/searchUser", api.SearchUser)
		//搜索帖子
		apiv1.GET("/searchArticle", v1.SearchArticle)

	}

	return r
}
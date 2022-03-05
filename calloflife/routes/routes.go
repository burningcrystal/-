package routes

import (
	v1 "calloflife/api/v1"
	"calloflife/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
/*
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
				c.JSON(http.StatusOK,gin.H{
					"msg":"ok",
				})
		})
	}*/
	routerV1 := r.Group("api/v1")
	{
		//User的路由结构
		routerV1.POST("user/add",v1.AddUser)
		routerV1.GET("users",v1.GetUserList)
		routerV1.PUT("user/:id",v1.EditUser)
		routerV1.DELETE("user/:id",v1.DeletUser)
		//category的路由结构
		routerV1.POST("category/add",v1.AddCategory)
		routerV1.GET("categorys",v1.GetCategoryList)
		routerV1.PUT("category/:id",v1.EditCategory)
		routerV1.DELETE("category/:id",v1.DeletCategory)
		//文章的路由结构
		routerV1.POST("article/add",v1.AddArticle)
		routerV1.GET("articles",v1.GetArticleList)  //获取文章列表
		routerV1.GET("articles/articlewithcate",v1.GetCateArt)  //获取分类下的所有文章
		routerV1.GET("article/info/:id",v1.GetArticleInfo) //获取单个文章信息
		routerV1.PUT("article/:id",v1.EditArticle)
		routerV1.DELETE("article/:id",v1.DeletArticle)
	}
	r.Run(utils.HttpPort)
}

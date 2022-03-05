package v1

import (
	"calloflife/model"
	"calloflife/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context){
	var data model.Article
	_ = c.ShouldBindJSON(&data)  //将接受到的json数据映射为data结构体
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//todo 查询分类下的所有文章
func GetCateArt(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber ,_ :=strconv.Atoi(c.Query("pagenumber"))
	if pageSize==0 {  //如果pagesize为0的话就不需要分页功能，按照gorm的规定传入-1取消limit的功能
		pageSize=-1
	}
	if pageNumber==0 {
		pageNumber=1
	}
	var data []model.Article
	cid,_ := strconv.Atoi(c.Query("cid"))
	code,data = model.GetCateArt(pageSize,pageNumber,cid)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"msg":errmsg.GetErrMsg(code),
	})
}
//todo 查询单个文章信息
func GetArticleInfo(c *gin.Context) {
	id,_:= strconv.Atoi(c.Param("id"))
	var data model.Article
	code,data = model.GetArticleInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data": data,
		"msg": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticleList(c *gin.Context){
	//接受两个前端传过来的query,注意gin的query函数返回的是字符串，要做一定的转换才能传入model的接口中
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber ,_ :=strconv.Atoi(c.Query("pagenumber"))
	if pageSize==0 {  //如果pagesize为0的话就不需要分页功能，按照gorm的规定传入-1取消limit的功能
		pageSize=-1
	}
	if pageNumber==0 {
		pageNumber=1
	}
	data,code := model.GetArticle(pageSize,pageNumber)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"msg":errmsg.GetErrMsg(code),
	})
}
//文章编辑
func EditArticle(c *gin.Context){
	var data  model.Article
	id,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArticle(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除文章
func DeletArticle(c *gin.Context){
	id,_:= strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"massage":errmsg.GetErrMsg(code),
	})
}
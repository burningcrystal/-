package v1

import (
	"calloflife/model"
	"calloflife/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


//todo:查询单个分类下的文章


//添加分类
func AddCategory(c *gin.Context){
	var data model.Category
	_ = c.ShouldBindJSON(&data)  //将接受到的json数据映射为data结构体
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED{
		code = errmsg.ERROR_CATEGORY_USED
	}
	if code == errmsg.ERROR{
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询分类列表
func GetCategoryList(c *gin.Context){
	//接受两个前端传过来的query,注意gin的query函数返回的是字符串，要做一定的转换才能传入model的接口中
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber ,_ :=strconv.Atoi(c.Query("pagenumber"))
	if pageSize==0 {  //如果pagesize为0的话就不需要分页功能，按照gorm的规定传入-1取消limit的功能
		pageSize=-1
	}
	if pageNumber==0 {
		pageNumber=1
	}
	code=errmsg.SUCCESS
	data := model.GetCategorys(pageSize,pageNumber)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"msg":errmsg.GetErrMsg(code),
	})
}
//分类编辑
func EditCategory(c *gin.Context){
	var data  model.Category
	id,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = model.EditCategory(id,&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}
	if code == errmsg.ERROR{
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除分类
func DeletCategory(c *gin.Context){
	id,_:= strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"massage":errmsg.GetErrMsg(code),
	})
}
package v1

import (
	"calloflife/model"
	"calloflife/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)
var code int
//查询用户是否已经存在
func UserExist(c *gin.Context){

}
//添加用户
func AddUser(c *gin.Context){
	var data model.User
	_ = c.ShouldBindJSON(&data)  //将接受到的json数据映射为data结构体
	code = model.CheckUser(data.Username)
	log.Println(data.Username,"+++++++++++++++++++++++++++++++++")
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询用户列表
func GetUserList(c *gin.Context){
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
	data := model.GetUsers(pageSize,pageNumber)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"msg":errmsg.GetErrMsg(code),
	})
}
//用户编辑
func EditUser(c *gin.Context){
	var data  model.User
	id,_:= strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = model.EditUser(id,&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除用户
func DeletUser(c *gin.Context){
	id,_:= strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"massage":errmsg.GetErrMsg(code),
	})
}
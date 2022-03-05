package model

import (
	"calloflife/utils/errmsg"
	"encoding/base64"
	"gorm.io/gorm"
	"log"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	//Avatar string `gorm:"type:" json:"avatar"`
	Role string `gorm:"type:int" json:"role"`
	//role用来设置用户的权限
}
//查询用户是否存在
func CheckUser(username string)int  {
	var users User  //声明一个结构体，也就是拉起一个到数据库user表的映射
	db.Select("id").Where("username = ?",username).First(&users)
	log.Println(users.ID,users.Username)
	if users.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(data *User)int  {
	//新增数据，并查看错误信息
	data.Password = ScryptPW(data.Password) //先对密码进行加密
	res := db.Create(data)
	log.Println("已写入",data.Username)
	if res.Error!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
//考虑到数据量大的可能性，要先做好分页代码的设定
//pagesize:数据库分页后一页有pagesize条数据
//pagenumber:分页后我们想看的第pagenumber条数据
func GetUsers(pageSize int , pageNumber int) []User {
	var users []User
	err := db.Offset((pageNumber-1)*pageSize).Limit(pageSize).Find(&users).Error
	if err!=nil&& err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//编辑用户
//先拿到用户ID，然后查询数据库，对数据库的内容做一下更新,UPDATE
func EditUser(id int,data *User)int{
	//注意gorm中用struct更新和用map更新的区别
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(user).Where("id = ?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
//也就是从数据库中删除对应的sample，并返回删除操作的状态
func DeleteUser(id int)int  {
	var user User
	err := db.Where("id = ?",id).Delete(&user).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPW(password string)string {
	const KeyLen=10
	salt := make([]byte,8)
	salt=[]byte{44,45,12,54,23,100,99,22}
	HashPw,err := scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err != nil {
		log.Println(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
	//这里注意各种调用手法
	//将加密后的密码写回数据库的方法之一：使用钩子函数；或者是在密码写入之前加密
	//开销有待测试
}

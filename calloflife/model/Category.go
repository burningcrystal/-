package model

import (
	"calloflife/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

//尽量简化一下分类的数据量。
type Category struct {
	ID int `gorm:"primary_Key;auto_increment;" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
//查询用户是否存在
func CheckCategory(name string)int  {
	var cate Category  //声明一个结构体，也就是拉起一个到数据库user表的映射
	db.Select("id").Where("name = ?",name).First(&cate)
	log.Println(cate.ID,cate.Name)
	if cate.ID > 0{
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateCategory(data *Category)int  {
	//新增数据，并查看错误信息
	res := db.Create(data)
	if res.Error!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
//考虑到数据量大的可能性，要先做好分页代码的设定
//pagesize:数据库分页后一页有pagesize条数据
//pagenumber:分页后我们想看的第pagenumber条数据
func GetCategorys(pageSize int , pageNumber int) []Category {
	var cate []Category
	err := db.Offset((pageNumber-1)*pageSize).Limit(pageSize).Find(&cate).Error
	if err!=nil&& err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//编辑用户
//先拿到用户ID，然后查询数据库，对数据库的内容做一下更新,UPDATE
func EditCategory(id int,data *Category)int{
	//注意gorm中用struct更新和用map更新的区别
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(cate).Where("id = ?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
//也就是从数据库中删除对应的sample，并返回删除操作的状态
func DeleteCategory(id int)int  {
	var cate Category
	err := db.Where("id = ?",id).Delete(&cate).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

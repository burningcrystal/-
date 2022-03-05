package model

import (
	"calloflife/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:cid"`
	Title string `gorm:"type:varchar(50);not null" json:"title"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img string `gorm:"type:varchar(100);not null" json:"img"`
	Cid int `gorm:"type:int;not null" json:"cid"` //即：categoryID
}


//新增文章
func CreateArticle(data *Article)int  {
	//新增数据，并查看错误信息
	res := db.Create(data)
	if res.Error!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//todo 查询分类下的所有文章
func GetCateArt(pageSize int , pageNumber int,cid int)(int,[]Article){
	var arti []Article
	err := db.Preload("Category").Offset((pageNumber-1)*pageSize).Limit(pageSize).Where("Cid=?",cid).Find(&arti).Error
	if err!=nil {
		return errmsg.ERROR_CATEGORY_NOT_EXIST,nil
	}
	return errmsg.SUCCESS,arti
}
//todo 查询单个文章
func GetArticleInfo(id int)(int,Article){
	var article Article
	err := db.Preload("Category").Where("id = ?",id).First(&article).Error
	if err != nil {
		return errmsg.ERROR_ARTI_NOT_EXIST,Article{}
	}
	return errmsg.SUCCESS,article
}

//查询文章列表
//考虑到数据量大的可能性，要先做好分页代码的设定
//pagesize:数据库分页后一页有pagesize条数据
//pagenumber:分页后我们想看的第pagenumber条数据
func GetArticle(pageSize int , pageNumber int) ([]Article,int) {
	var arti []Article

	err := db.Preload("Category").Offset((pageNumber-1)*pageSize).Limit(pageSize).Find(&arti).Error
	if err!=nil&& err != gorm.ErrRecordNotFound {
		return nil,errmsg.ERROR
	}
	return arti,errmsg.SUCCESS
}

//编辑用户
//先拿到用户ID，然后查询数据库，对数据库的内容做一下更新,UPDATE
func EditArticle(id int,data *Article)int{
	//注意gorm中用struct更新和用map更新的区别
	var arti Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["description"] = data.Description
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["cid"] = data.Cid


	err = db.Model(arti).Where("id = ?",id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
//也就是从数据库中删除对应的sample，并返回删除操作的状态
func DeleteArticle(id int)int  {
	var arti Article
	err := db.Where("id = ?",id).Delete(&arti).Error
	if err!=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
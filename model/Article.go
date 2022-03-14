package model

import (
	"gin_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	if err := db.Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询文章下的所有文章

// todo 查询单个文章

// todo 查询文章列表
func GetArts(pageSize int, pageNum int) []Article {
	var arts []Article
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return arts
}

// EditArt 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

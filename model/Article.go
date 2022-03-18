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

// GetCateArts 查询分类下的所有文章
func GetCateArts(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", id).Count(&total) // preload is not allowed when count is used
	if err != nil || len(cateArtList) == 0 {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArts 查询文章列表
func GetArts(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	db.Model(&articleList).Count(&total) // preload is not allowed when count is used
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
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

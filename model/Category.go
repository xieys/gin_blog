package model

import (
	"gin_blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCate 查询分类是否存在
func CheckCate(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	if err := db.Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCates 查询分类列表
func GetCates(pageSize int, pageNum int) ([]Category, int64) {
	var cates []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// EditCate 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

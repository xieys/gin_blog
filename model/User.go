package model

import (
	"encoding/base64"
	"gin_blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	if err := db.Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户

// 删除用户

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	keyLen := 10
	salt := make([]byte, 8)
	salt = []byte{32, 5, 8, 22, 123, 18, 56, 14}

	hashPw, err := scrypt.Key([]byte(password), salt, 1<<11, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(hashPw)
	return fpw
}

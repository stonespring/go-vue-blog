package models

import "gin.blog.com/db"

type User struct {
	Username     string     `gorm:"username" json:"username"`
	Password     string     `gorm:"password" json:"password"`
	ID           int 		`gorm:"primary_key" json:"id"`
	PIC          string 	`gorm:"pic" json:"pic"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//验证用户
func CheckAuth(username, password string) bool {
	var user User
	db.GetDb().Table("users").Select("id").Where(User{Username : username, Password : password}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}


func GetUserOne(userName string) (u User) {

	db.GetDb().Table("users").
		Where("username = ?", userName).
		Find(&u)
	return
}



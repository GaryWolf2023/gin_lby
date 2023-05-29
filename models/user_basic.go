package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Role       int       `json:"role"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTiem time.Time `json:"logout_time"`
}

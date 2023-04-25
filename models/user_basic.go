package models

import "time"

type UserBasic struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Role       int       `json:"role"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTiem time.Time `json:"logout_time"`
}

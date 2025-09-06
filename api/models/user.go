package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	PassHash string `json:"passhash" gorm:"not null"`
}

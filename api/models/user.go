package models

type User struct {
	UuId    string  `json:"uuid" gorm:"primaryKey"`
	Email   string  `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Passwrd string  `json:"passwrd" gorm:"size:256;not null"`
	SessId  *string `json:"sessid" gorm:"size:256"`
}

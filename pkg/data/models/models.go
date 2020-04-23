package models

import "time"

// BaseModel definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`,
// which could be embedded in your models
//    type User struct {
//      BaseModelModel
//    }
type BaseModel struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	DeletedAt time.Time `json:"deletedat" sql:"index"`
}

type User struct {
	BaseModel
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
	UserName    string `json:"username"`
}

type Comment struct {
	BaseModel
	Content string `json:"content"`
}

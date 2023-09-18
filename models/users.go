package models

type Student struct {
	Id    int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique not null"`
	Grade int    `json:"grade" gorm:"not null"`
}
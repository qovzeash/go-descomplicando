package models

type User struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

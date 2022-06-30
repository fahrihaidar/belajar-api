package models

type User struct {
	Id       int
	Name     string
	Email    string `gorm:"unique_index"`
	Password []byte
}

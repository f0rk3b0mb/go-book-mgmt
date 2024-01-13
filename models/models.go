package model

type Book struct {
	ID     uint
	Name   string
	Author string
}

type Users struct {
	ID       uint
	Username string
	Email    string `gorm:"unique"`
	Password string
}

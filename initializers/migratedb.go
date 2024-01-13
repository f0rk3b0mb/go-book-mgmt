package initializer

import model "github.com/f0rk3b0mb/go-book-mgmt/models"

func Migratedb() {
	var book model.Book
	var user model.Users
	DB.AutoMigrate(&book)
	DB.AutoMigrate(&user)
}

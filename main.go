package main

import (
	initializer "github.com/f0rk3b0mb/go-book-mgmt/initializers"
	middleware "github.com/f0rk3b0mb/go-book-mgmt/middleware"
	routes "github.com/f0rk3b0mb/go-book-mgmt/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.GetenvironmentVars()
	initializer.Connectdb()
	initializer.Migratedb()
}

func main() {
	r := gin.Default()
	r.GET("/view", middleware.Validate, routes.View)
	r.POST("/add", middleware.Validate, routes.Addbook)
	r.GET("/search/:id", middleware.Validate, routes.Searchbook)
	r.DELETE("/delete/:id", middleware.Validate, routes.Delete)
	r.POST("/login", routes.Login)
	r.POST("register", routes.Register)

	r.Run()
}

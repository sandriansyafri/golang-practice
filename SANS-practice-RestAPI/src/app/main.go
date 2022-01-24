package main

import (
	"sans-api-pustaka-api/book"
	"sans-api-pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/go_api-project_pustaka?charset=utf8mb4&parseTime=True&loc=Local" // TODO configuration DB connection
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&book.Book{}) //TODO migrate table

	// layering
	RepositoryBook := book.NewRepositoryBook(db)
	ServiceBook := book.NewServiceBook(RepositoryBook)
	BookHandler := handler.NewBookHandler(ServiceBook)

	r := gin.Default() //create router

	v1 := r.Group("/v1") // group routes

	// routes
	v1.GET("/", BookHandler.RootHandler)
	v1.GET("/books", BookHandler.GetBooks)
	v1.GET("/books/:id", BookHandler.GetBook)
	v1.POST("/books", BookHandler.CreateBook)
	v1.PUT("/books/:id", BookHandler.UpdateBook)
	v1.DELETE("/books/:id", BookHandler.DeleteBook)

	r.Run(":8000") //run web server
}

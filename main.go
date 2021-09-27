package main

import (
	"book-app/auth"
	"book-app/config"
	"book-app/handler"
	"book-app/middleware"
	"book-app/repository"
	"book-app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware()) // ! Allow cors
	
	//! Auth
	authService := auth.NewService()

	//! Users
	userRepository := repository.NewRepositoryUser(config.ConfigDB())
	userService := service.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	userGroup := router.Group("/users")
	userGroup.POST("/", userHandler.RegisterUser)
	userGroup.POST("/login", userHandler.LoginUser)
	userGroup.GET("/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)
	userGroup.PUT("/", middleware.AuthMiddleware(authService, userService), userHandler.UpdateUser)

	bookRepository := repository.NewRepositoryBook(config.ConfigDB())
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	bookGroup := router.Group("/books")
	bookGroup.GET("/", bookHandler.GetBooks)
	bookGroup.GET("/:id", bookHandler.GetBook)
	bookGroup.PUT("/:id", middleware.AuthMiddleware(authService, userService), bookHandler.UpdateBook)
	bookGroup.POST("/", middleware.AuthMiddleware(authService, userService), bookHandler.CreateBook)
	bookGroup.DELETE("/:id", middleware.AuthMiddleware(authService, userService), bookHandler.DeleteBook)

	router.Run() //! default PORT 8080
}

/*
	! Kekurangan:
	? Validasi masih belum bisa jika ada usernam yang sudah terdaftar
	? Cara token ada masa expired nya
	? Refresh token
*/

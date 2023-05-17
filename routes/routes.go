package routes

import (
	"log"

	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/middlewares"
	service "github.com/dhxmo/shop-stop-go/services"
	"github.com/gin-gonic/gin"
)

func Routes(e *gin.Engine) {
	auth := e.Group("auth")
	{
		userService := service.NewUserSvc()
		auth.POST("/register", userService.Register)
		auth.POST("/login", userService.Login)
		auth.GET("/users/:uuid", userService.GetUserByID)

	}

	v1 := e.Group("api/v1")
	v1.Use(middlewares.JWT())
	v1.Use(middlewares.ErrorHandler())

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if config.Enable {
		v1.Use(middlewares.Cached())
	}
	{
		productService := service.NewProductSvc()
		v1.GET("/products", productService.GetProducts)
		v1.GET("/products/:uuid", productService.GetProductByID)
		v1.POST("/product", productService.CreateProduct)
		v1.PUT("/product/:uuid", productService.UpdateProduct)

		categoryService := service.NewCategorySvc()
		v1.GET("/categories", categoryService.GetCategories)
		v1.POST("/categories", categoryService.CreateCategory)
		v1.GET("/categories/:uuid", categoryService.GetCategoryByID)
		v1.GET("/categories/:uuid/products", productService.GetProductByCategory)
		v1.PUT("/categories/:uuid", categoryService.UpdateCategory)

	}
}

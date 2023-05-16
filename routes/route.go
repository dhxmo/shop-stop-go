package routes

import (
	"github.com/dhxmo/shop-stop-go/middlewares"
	service "github.com/dhxmo/shop-stop-go/services"
	"github.com/gin-gonic/gin"
)

func API(e *gin.Engine) {
	userService := service.NewUserSvc()
	e.POST("/register", userService.Register)
	e.POST("/login", userService.Login)

	v1 := e.Group("api/v1")
	v1.Use(middlewares.JWT())
	v1.Use(middlewares.ErrorHandler())

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

		v1.GET("/users/:uuid", userService.GetUserByID)
	}
}

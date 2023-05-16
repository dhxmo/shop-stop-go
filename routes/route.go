package routes

import (
	service "github.com/dhxmo/shop-stop-go/services"
	"github.com/gin-gonic/gin"
)

func API(e *gin.Engine) {
	v1 := e.Group("api/v1")
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

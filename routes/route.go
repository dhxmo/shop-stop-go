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
	}
}

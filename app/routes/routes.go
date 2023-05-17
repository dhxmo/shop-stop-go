package routes

import (
	"log"

	"github.com/dhxmo/shop-stop-go/app/middlewares"
	service "github.com/dhxmo/shop-stop-go/app/services"
	"github.com/dhxmo/shop-stop-go/config"
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

		quantityService := service.NewQuantitySvc()
		v1.GET("/quantities", quantityService.GetQuantities)
		v1.POST("/quantities", quantityService.CreateQuantity)
		v1.GET("/quantities/:uuid", quantityService.GetQuantityByID)
		v1.PUT("/quantities/:uuid", quantityService.UpdateQuantity)

		checkoutService := service.NewCheckoutSvc()
		v1.GET("/checkout", checkoutService.GetCheckouts)
		v1.POST("/checkout", checkoutService.CreateCheckout)
		v1.GET("/checkout/:uuid", checkoutService.GetCheckoutByID)
		v1.PUT("/checkout/:uuid", checkoutService.UpdateCheckout)

		orderService := service.NewOrderSvc()
		v1.GET("/orders", orderService.GetOrders)
		v1.POST("/orders", orderService.CreateOrder)
		v1.GET("/orders/:uuid", orderService.GetOrderByID)
		v1.PUT("/orders/:uuid", orderService.UpdateOrder)

	}
}

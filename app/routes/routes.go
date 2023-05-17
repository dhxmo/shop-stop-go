package routes

import (
	"log"

	controllers "github.com/dhxmo/shop-stop-go/app/controllers"
	middlewares "github.com/dhxmo/shop-stop-go/app/middlewares"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Routes(e *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		category *controllers.Category,
		product *controllers.Product,
		quantity *controllers.Quantity,
		checkout *controllers.Checkout,
		order *controllers.Order,
		user *controllers.User,
	) error {
		auth := e.Group("auth")
		{
			auth.POST("/register", user.Register)
			auth.POST("/login", user.Login)
			auth.GET("/users/:uuid", user.GetUserByID)

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
			v1.GET("/products", product.GetProducts)
			v1.GET("/products/:uuid", product.GetProductByID)
			v1.POST("/product", product.CreateProduct)
			v1.PUT("/product/:uuid", product.UpdateProduct)
		}
		{
			v1.GET("/categories", category.GetCategories)
			v1.POST("/categories", category.CreateCategory)
			v1.GET("/categories/:uuid", category.GetCategoryByID)
			v1.GET("/categories/:uuid/products ", product.GetProductByCategory)
			v1.PUT("/categories/:uuid", category.UpdateCategory)
		}
		{
			v1.GET("/quantities", quantity.GetQuantities)
			v1.POST("/quantities", quantity.CreateQuantity)
			v1.GET("/quantities/:uuid", quantity.GetQuantityByID)
			v1.PUT("/quantities/:uuid", quantity.UpdateQuantity)
		}
		{
			v1.GET("/checkout", checkout.GetCheckouts)
			v1.POST("/checkout", checkout.CreateCheckout)
			v1.GET("/checkout/:uuid", checkout.GetCheckoutByID)
			v1.PUT("/checkout/:uuid", checkout.UpdateCheckout)
		}
		{
			v1.GET("/orders", order.GetOrders)
			v1.POST("/orders", order.CreateOrder)
			v1.GET("/orders/:uuid", order.GetOrderByID)
			v1.PUT("/orders/:uuid", order.UpdateOrder)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return err
}

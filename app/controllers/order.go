package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
)

type Order struct {
	service services.OrderService
}

func NewOrderController(service services.OrderService) *Order {
	return &Order{service: service}
}

func (os *Order) GetOrders(c *gin.Context) {
	ctx := c.Request.Context()
	orders, err := os.service.GetOrders(ctx)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(orders, "ok", ""))
}

func (os *Order) GetOrderByID(c *gin.Context) {
	orderID := c.Param("uuid")
	ctx := c.Request.Context()

	order, err := os.service.GetOrderByID(ctx, orderID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(order, "ok", ""))
}

func (os *Order) CreateOrder(c *gin.Context) {
	var req models.OrderRequest

	if err := c.Bind(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	ctx := c.Request.Context()

	order, err := os.service.CreateOrder(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(order, "OK", ""))
}

func (os *Order) UpdateOrder(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	orders, err := os.service.UpdateOrder(ctx, uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(orders, "OK", ""))
}

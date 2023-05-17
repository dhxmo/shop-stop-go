package services

import (
	"net/http"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type OrderService interface {
	GetOrders(c *gin.Context)
	GetOrderByID(c *gin.Context)
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
}

type OrderSvc struct {
	repo repositories.OrderRepository
}

func NewOrderSvc() OrderService {
	return &OrderSvc{repo: repositories.NewOrderRepository()}
}

func (os *OrderSvc) GetOrders(c *gin.Context) {
	orders, err := os.repo.GetOrders()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(orders, "ok", ""))
}

func (os *OrderSvc) GetOrderByID(c *gin.Context) {
	orderID := c.Param("uuid")

	order, err := os.repo.GetOrderByID(orderID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(order, "ok", ""))
}

func (os *OrderSvc) CreateOrder(c *gin.Context) {
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

	order, err := os.repo.CreateOrder(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(order, "OK", ""))
}

func (os *OrderSvc) UpdateOrder(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := os.repo.UpdateOrder(uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(orders, "OK", ""))
}

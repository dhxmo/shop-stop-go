package services

import (
	"net/http"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CheckoutService interface {
	GetCheckouts(c *gin.Context)
	GetCheckoutByID(c *gin.Context)
	CreateCheckout(c *gin.Context)
	UpdateCheckout(c *gin.Context)
}

type CheckoutSvc struct {
	repo repositories.CheckoutOrderRepository
}

func NewCheckoutSvc() CheckoutService {
	return &CheckoutSvc{repo: repositories.NewCheckoutOrderRepository()}
}

func (cs *CheckoutSvc) GetCheckouts(c *gin.Context) {
	checkoutOrders, err := cs.repo.GetCheckoutOrders()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(checkoutOrders, "ok", ""))
}

func (cs *CheckoutSvc) GetCheckoutByID(c *gin.Context) {
	checkoutOrdersID := c.Param("uuid")

	checkoutOrders, err := cs.repo.GetCheckoutOrderByID(checkoutOrdersID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(checkoutOrders, "ok", ""))
}

func (cs *CheckoutSvc) CreateCheckout(c *gin.Context) {
	var req models.CheckoutOrderBodyRequest

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

	checkout, err := cs.repo.CreateCheckoutOrder(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(checkout, "OK", ""))
}

func (cs *CheckoutSvc) UpdateCheckout(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.CheckoutOrderBodyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkout, err := cs.repo.UpdateCheckoutOrder(uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(checkout, "OK", ""))
}

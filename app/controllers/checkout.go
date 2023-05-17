package controllers

import (
	"net/http"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Checkout struct {
	service services.CheckoutService
}

func NewCheckoutController(service services.CheckoutService) *Checkout {
	return &Checkout{
		service: service,
	}
}

func (cs *Checkout) GetCheckouts(c *gin.Context) {
	ctx := c.Request.Context()
	checkoutOrders, err := cs.service.GetCheckouts(ctx)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(checkoutOrders, "ok", ""))
}

func (cs *Checkout) GetCheckoutByID(c *gin.Context) {
	checkoutOrdersID := c.Param("uuid")
	ctx := c.Request.Context()

	checkoutOrders, err := cs.service.GetCheckoutByID(ctx, checkoutOrdersID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(checkoutOrders, "ok", ""))
}

func (cs *Checkout) CreateCheckout(c *gin.Context) {
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
	ctx := c.Request.Context()

	checkout, err := cs.service.CreateCheckout(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(checkout, "OK", ""))
}

func (cs *Checkout) UpdateCheckout(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.CheckoutOrderBodyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()

	checkout, err := cs.service.UpdateCheckout(ctx, uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(checkout, "OK", ""))
}

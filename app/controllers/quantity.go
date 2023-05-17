package controllers

import (
	"net/http"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Quantity struct {
	service services.QuantityService
}

func NewQuantityController(service services.QuantityService) *Quantity {
	return &Quantity{
		service: service,
	}
}

func (q *Quantity) GetQuantities(c *gin.Context) {
	ctx := c.Request.Context()

	quantities, err := q.service.GetQuantities(ctx)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(quantities, "ok", ""))
}

func (q *Quantity) GetQuantityByID(c *gin.Context) {
	quantitesID := c.Param("uuid")

	ctx := c.Request.Context()
	quantity, err := q.service.GetQuantityByID(ctx, quantitesID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(quantity, "ok", ""))
}

func (q *Quantity) CreateQuantity(c *gin.Context) {
	var req models.QuantityBodyRequest

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
	quantity, err := q.service.CreateQuantity(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(quantity, "OK", ""))
}

func (q *Quantity) UpdateQuantity(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.QuantityBodyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	quantity, err := q.service.UpdateQuantity(ctx, uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(quantity, "OK", ""))
}

package service

import (
	"net/http"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type QuantityService interface {
	GetQuantities(c *gin.Context)
	GetQuantityByID(c *gin.Context)
	CreateQuantity(c *gin.Context)
	UpdateQuantity(c *gin.Context)
}

type QuantitySvc struct {
	repo repositories.QuantityRepository
}

func NewQuantitySvc() QuantityService {
	return &QuantitySvc{repo: repositories.NewQuantityRepository()}
}

func (qs *QuantitySvc) GetQuantities(c *gin.Context) {
	quantities, err := qs.repo.GetQuantities()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(quantities, "ok", ""))
}

func (qs *QuantitySvc) GetQuantityByID(c *gin.Context) {
	quantitesID := c.Param("uuid")

	quantity, err := qs.repo.GetQuantityByID(quantitesID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(quantity, "ok", ""))
}

func (qs *QuantitySvc) CreateQuantity(c *gin.Context) {
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

	quantity, err := qs.repo.CreateQuantity(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(quantity, "OK", ""))
}

func (qs *QuantitySvc) UpdateQuantity(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.QuantityBodyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quantity, err := qs.repo.UpdateQuantity(uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(quantity, "OK", ""))
}

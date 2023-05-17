package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
)

type Product struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *Product {
	return &Product{service: service}
}

func (product *Product) GetProducts(c *gin.Context) {
	var params models.ProductQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := product.service.GetProducts(ctx, params)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res []models.ProductResponse
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

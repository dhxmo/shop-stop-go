package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

func (product *Product) GetProductByID(c *gin.Context) {
	productID := c.Param("uuid")

	ctx := c.Request.Context()
	prod, err := product.service.GetProductByID(ctx, productID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(prod, "ok", ""))
}

func (product *Product) CreateProduct(c *gin.Context) {
	var req models.ProductRequest

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
	prod, err := product.service.CreateProduct(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(prod, "OK", ""))
}

func (product *Product) UpdateProduct(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	products, err := product.service.UpdateProduct(ctx, uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(products, "OK", ""))
}

func (product *Product) GetProductByCategory(c *gin.Context) {
	categoryUUID := c.Param("uuid")
	activeParam := c.Query("active")
	active := true
	if activeParam == "false" {
		active = false
	}

	ctx := c.Request.Context()
	products, err := product.service.GetProductByCategory(ctx, categoryUUID, active)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(products, "OK", ""))
}

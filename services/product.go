package service

import (
	"net/http"

	"github.com/dhxmo/shop-stop-go/models"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/dhxmo/shop-stop-go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProductService interface {
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	GetProductByCategory(c *gin.Context)
}

type ProductSvc struct {
	repo repositories.ProductRepository
}

func NewProductSvc() ProductService {
	return &ProductSvc{repo: repositories.NewProductRepository()}
}

func (ps *ProductSvc) GetProducts(c *gin.Context) {
	products, err := ps.repo.GetProducts()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(products, "ok", ""))
}

func (ps *ProductSvc) GetProductByID(c *gin.Context) {
	productID := c.Param("uuid")

	product, err := ps.repo.GetProductByID(productID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(product, "ok", ""))
}

func (ps *ProductSvc) CreateProduct(c *gin.Context) {
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

	product, err := ps.repo.CreateProduct(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(product, "OK", ""))
}

func (ps *ProductSvc) UpdateProduct(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := ps.repo.UpdateProduct(uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(products, "OK", ""))
}

func (ps *ProductSvc) GetProductByCategory(c *gin.Context) {
	categoryUUID := c.Param("uuid")
	activeParam := c.Query("active")
	active := true
	if activeParam == "false" {
		active = false
	}

	products, err := ps.repo.GetProductByCategory(categoryUUID, active)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(products, "OK", ""))
}

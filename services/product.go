package service

import (
	"log"
	"net/http"

	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/dhxmo/shop-stop-go/repositories"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
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
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(products, "ok", ""))
}

func (ps *ProductSvc) GetProductByID(c *gin.Context) {
	productID := c.Param("uuid")

	product, err := ps.repo.GetProductByID(productID)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(product, "ok", ""))
}

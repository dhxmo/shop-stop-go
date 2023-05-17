package controllers

import (
	"net/http"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Category struct {
	Service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *Category {
	return &Category{
		Service: service,
	}
}

func (categ *Category) GetCategories(c *gin.Context) {
	var reqQuery models.CategoryQueryRequest
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := categ.Service.GetCategories(ctx, reqQuery)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res []models.CategoryResponse
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

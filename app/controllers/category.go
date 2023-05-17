package controllers

import (
	"net/http"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type Category struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *Category {
	return &Category{
		service: service,
	}
}

func (category *Category) GetCategories(c *gin.Context) {
	var reqQuery models.CategoryQueryRequest
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	rs, err := category.service.GetCategories(ctx, reqQuery)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res []models.CategoryResponse
	copier.Copy(&res, &rs)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (category *Category) GetCategoryByID(c *gin.Context) {
	categoryId := c.Param("uuid")

	ctx := c.Request.Context()
	categ, err := category.service.GetCategoryByID(ctx, categoryId)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.CategoryResponse
	copier.Copy(&res, &categ)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (category *Category) CreateCategory(c *gin.Context) {
	var item models.CategoryRequest
	if err := c.Bind(&item); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(item)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	ctx := c.Request.Context()
	categories, err := category.service.CreateCategory(ctx, &item)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.CategoryResponse
	copier.Copy(&res, &categories)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (category *Category) UpdateCategory(c *gin.Context) {
	uuid := c.Param("uuid")
	var item models.CategoryRequest
	if err := c.ShouldBind(&item); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	categories, err := category.service.UpdateCategory(ctx, uuid, &item)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.CategoryResponse
	copier.Copy(&res, &categories)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

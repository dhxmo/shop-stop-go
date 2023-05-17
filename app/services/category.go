package service

import (
	"net/http"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryService interface {
	GetCategories(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
}

type CategorySvc struct {
	repo repositories.CategoryRepository
}

func NewCategorySvc() CategoryService {
	return &CategorySvc{repo: repositories.NewCategoryRepository()}
}

func (ps *CategorySvc) GetCategories(c *gin.Context) {
	categories, err := ps.repo.GetCategories()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(categories, "ok", ""))
}

func (ps *CategorySvc) GetCategoryByID(c *gin.Context) {
	categoryID := c.Param("uuid")

	category, err := ps.repo.GetCategoryByID(categoryID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Response(category, "ok", ""))
}

func (ps *CategorySvc) CreateCategory(c *gin.Context) {
	var req models.CategoryRequest

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

	category, err := ps.repo.CreateCategory(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(category, "OK", ""))
}

func (ps *CategorySvc) UpdateCategory(c *gin.Context) {
	uuid := c.Param("uuid")
	var req models.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := ps.repo.UpdateCategory(uuid, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Response(category, "OK", ""))
}

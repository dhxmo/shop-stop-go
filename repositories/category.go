package repositories

import (
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/models"
	"github.com/jinzhu/copier"
)

type CategoryRepository interface {
	GetCategories() (*[]models.CategoryResponse, error)
	GetCategoryByID(uuid string) (*models.CategoryResponse, error)
	CreateCategory(req *models.CategoryRequest) (*models.CategoryResponse, error)
	UpdateCategory(uuid string, req *models.CategoryRequest) (*models.CategoryResponse, error)
}

type CategoryRepo struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepo{}
}

func (r *CategoryRepo) GetCategories() (*[]models.CategoryResponse, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return &[]models.CategoryResponse{}, nil
	}

	var res []models.CategoryResponse
	copier.Copy(&res, &categories)

	return &res, nil
}

func (r *CategoryRepo) GetCategoryByID(uuid string) (*models.CategoryResponse, error) {
	var category models.Category
	if err := config.DB.Where("uuid = ?", uuid).Find(&category).Error; err != nil {
		return nil, err
	}

	if category.UUID == "" {
		return nil, nil
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

func (r *CategoryRepo) CreateCategory(req *models.CategoryRequest) (*models.CategoryResponse, error) {
	var category models.Category

	copier.Copy(&category, &req)
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

func (r *CategoryRepo) UpdateCategory(uuid string, req *models.CategoryRequest) (*models.CategoryResponse, error) {
	var category models.Category

	if err := config.DB.Where("uuid = ?", uuid).First(&category).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	category.Active = req.Active

	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

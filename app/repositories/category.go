package repositories

import (
	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	GetCategories(query models.CategoryQueryRequest) (*[]models.CategoryResponse, error)
	GetCategoryByID(uuid string) (*models.CategoryResponse, error)
	CreateCategory(req *models.CategoryRequest) (*models.CategoryResponse, error)
	UpdateCategory(uuid string, req *models.CategoryRequest) (*models.CategoryResponse, error)
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepo{db: config.DB}
}

func (cr *CategoryRepo) GetCategories(query models.CategoryQueryRequest) (*[]models.CategoryResponse, error) {
	var categories []models.Category
	if err := cr.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return &[]models.CategoryResponse{}, nil
	}

	var res []models.CategoryResponse
	copier.Copy(&res, &categories)

	return &res, nil
}

func (cr *CategoryRepo) GetCategoryByID(uuid string) (*models.CategoryResponse, error) {
	var category models.Category
	if err := cr.db.Where("uuid = ?", uuid).Find(&category).Error; err != nil {
		return nil, err
	}

	if category.UUID == "" {
		return nil, nil
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

func (cr *CategoryRepo) CreateCategory(req *models.CategoryRequest) (*models.CategoryResponse, error) {
	var category models.Category

	copier.Copy(&category, &req)
	if err := cr.db.Create(&category).Error; err != nil {
		return nil, err
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

func (cr *CategoryRepo) UpdateCategory(uuid string, req *models.CategoryRequest) (*models.CategoryResponse, error) {
	var category models.Category

	if err := cr.db.Where("uuid = ?", uuid).First(&category).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	category.Active = req.Active

	if err := cr.db.Save(&category).Error; err != nil {
		return nil, err
	}

	var res models.CategoryResponse
	copier.Copy(&res, &category)

	return &res, nil
}

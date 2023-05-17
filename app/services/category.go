package services

import (
	"context"
	"log"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type CategoryService interface {
	GetCategories(ctx context.Context, query models.CategoryQueryRequest) (*[]models.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, categoryID string) (*models.CategoryResponse, error)
	CreateCategory(ctx context.Context, query *models.CategoryRequest) (*models.CategoryResponse, error)
	UpdateCategory(ctx context.Context, uuid string, query *models.CategoryRequest) (*models.CategoryResponse, error)
}

type CategorySvc struct {
	repo repositories.CategoryRepository
}

func NewCategorySvc() CategoryService {
	return &CategorySvc{repo: repositories.NewCategoryRepository()}
}

func (ps *CategorySvc) GetCategories(ctx context.Context, query models.CategoryQueryRequest) (*[]models.CategoryResponse, error) {
	categories, err := ps.repo.GetCategories(query)
	if err != nil {
		log.Println("Failed to get categories: ", err)
		return nil, err
	}

	return categories, nil

}

func (ps *CategorySvc) GetCategoryByID(ctx context.Context, categoryID string) (*models.CategoryResponse, error) {
	category, err := ps.repo.GetCategoryByID(categoryID)
	if err != nil {
		log.Println("Failed to get categories: ", err)
		return nil, err
	}
	return category, nil
}

func (ps *CategorySvc) CreateCategory(ctx context.Context, query *models.CategoryRequest) (*models.CategoryResponse, error) {
	category, err := ps.repo.CreateCategory(query)
	if err != nil {
		log.Println("Failed to create category", err.Error())
		return nil, err
	}

	return category, nil
}

func (ps *CategorySvc) UpdateCategory(ctx context.Context, uuid string, query *models.CategoryRequest) (*models.CategoryResponse, error) {
	category, err := ps.repo.UpdateCategory(uuid, query)
	if err != nil {
		log.Println("Failed to update category", err.Error())
		return nil, err
	}

	return category, nil
}

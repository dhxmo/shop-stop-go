package services

import (
	"context"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type ProductService interface {
	GetProducts(ctx context.Context, params models.ProductQueryParams) (*[]models.ProductResponse, error)
	GetProductByID(ctx context.Context, productID string) (*models.ProductResponse, error)
	CreateProduct(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error)
	UpdateProduct(ctx context.Context, uuid string, req *models.ProductRequest) (*models.ProductResponse, error)
	GetProductByCategory(ctx context.Context, categoryUUID string, active bool) (*[]models.ProductResponse, error)
}

type ProductSvc struct {
	repo repositories.ProductRepository
}

func NewProductSvc() ProductService {
	return &ProductSvc{repo: repositories.NewProductRepository()}
}

func (ps *ProductSvc) GetProducts(ctx context.Context, params models.ProductQueryParams) (*[]models.ProductResponse, error) {
	products, err := ps.repo.GetProducts()
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return products, nil
}

func (ps *ProductSvc) GetProductByID(ctx context.Context, productID string) (*models.ProductResponse, error) {
	product, err := ps.repo.GetProductByID(productID)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return product, nil
}

func (ps *ProductSvc) CreateProduct(ctx context.Context, req *models.ProductRequest) (*models.ProductResponse, error) {
	product, err := ps.repo.CreateProduct(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return product, nil
}

func (ps *ProductSvc) UpdateProduct(ctx context.Context, uuid string, req *models.ProductRequest) (*models.ProductResponse, error) {
	prod, err := ps.repo.UpdateProduct(uuid, req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return prod, nil

}

func (ps *ProductSvc) GetProductByCategory(ctx context.Context, categoryUUID string, active bool) (*[]models.ProductResponse, error) {
	products, err := ps.repo.GetProductByCategory(categoryUUID, active)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return products, nil

}

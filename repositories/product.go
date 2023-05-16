package repositories

import (
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/models"
	"github.com/jinzhu/copier"
)

type ProductRepository interface {
	GetProducts() (*[]models.ProductResponse, error)
	GetProductByID(uuid string) (*models.ProductResponse, error)
	CreateProduct(req *models.ProductRequest) (*models.ProductResponse, error)
	UpdateProduct(uuid string, req *models.ProductRequest) (*models.ProductResponse, error)
}

type ProductRepo struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepo{}
}

func (r *ProductRepo) GetProducts() (*[]models.ProductResponse, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return &[]models.ProductResponse{}, nil
	}

	var res []models.ProductResponse
	copier.Copy(&res, &products)

	return &res, nil
}

func (r *ProductRepo) GetProductByID(uuid string) (*models.ProductResponse, error) {
	var product models.Product
	if err := config.DB.Where("uuid = ?", uuid).Find(&product).Error; err != nil {
		return nil, err
	}

	if product.UUID == "" {
		return nil, nil
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

func (r *ProductRepo) CreateProduct(req *models.ProductRequest) (*models.ProductResponse, error) {
	var product models.Product

	copier.Copy(&product, &req)
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

func (r *ProductRepo) UpdateProduct(uuid string, req *models.ProductRequest) (*models.ProductResponse, error) {
	var product models.Product

	if err := config.DB.Where("uuid = ?", uuid).First(&product).Error; err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

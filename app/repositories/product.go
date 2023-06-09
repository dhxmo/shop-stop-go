package repositories

import (
	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetProducts() (*[]models.ProductResponse, error)
	GetProductByID(uuid string) (*models.ProductResponse, error)
	CreateProduct(req *models.ProductRequest) (*models.ProductResponse, error)
	UpdateProduct(uuid string, req *models.ProductRequest) (*models.ProductResponse, error)
	GetProductByCategory(uuid string, active bool) (*[]models.ProductResponse, error)
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &ProductRepo{db: config.DB}
}

func (pr *ProductRepo) GetProducts() (*[]models.ProductResponse, error) {
	var products []models.Product
	if err := pr.db.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return &[]models.ProductResponse{}, nil
	}

	var res []models.ProductResponse
	copier.Copy(&res, &products)

	return &res, nil
}

func (pr *ProductRepo) GetProductByCategory(categoryUUID string, active bool) (*[]models.ProductResponse, error) {
	var products []models.Product
	if err := pr.db.Where("active = ? AND category_uuid = ?", active, categoryUUID).Find(&products).Error; err != nil {
		return nil, err
	}

	var res []models.ProductResponse
	copier.Copy(&res, &products)
	return &res, nil
}

func (pr *ProductRepo) GetProductByID(uuid string) (*models.ProductResponse, error) {
	var product models.Product
	if err := pr.db.Where("uuid = ?", uuid).Find(&product).Error; err != nil {
		return nil, err
	}

	if product.UUID == "" {
		return nil, nil
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

func (pr *ProductRepo) CreateProduct(req *models.ProductRequest) (*models.ProductResponse, error) {
	var product models.Product

	copier.Copy(&product, &req)
	if err := pr.db.Create(&product).Error; err != nil {
		return nil, err
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

func (pr *ProductRepo) UpdateProduct(uuid string, req *models.ProductRequest) (*models.ProductResponse, error) {
	var product models.Product

	if err := pr.db.Where("uuid = ?", uuid).First(&product).Error; err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description

	if err := pr.db.Save(&product).Error; err != nil {
		return nil, err
	}

	var res models.ProductResponse
	copier.Copy(&res, &product)

	return &res, nil
}

package repositories

import (
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/models"
)

type ProductRepository interface {
	GetProducts() (*[]models.Product, error)
	GetProductByID(uuid string) (*models.Product, error)
}

type ProductRepo struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepo{}
}

func (r *ProductRepo) GetProducts() (*[]models.Product, error) {
	var products []models.Product
	if config.DB.Find(&products).RecordNotFound() {
		return nil, nil
	}

	return &products, nil
}

func (r *ProductRepo) GetProductByID(uuid string) (*models.Product, error) {
	var product models.Product
	if config.DB.Where("uuid = ?", uuid).Find(&product).RecordNotFound() {
		return nil, nil
	}
	return &product, nil
}

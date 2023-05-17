package repositories

import (
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/models"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	GetOrders() (*[]models.OrderResponse, error)
	GetOrderByID(uuid string) (*models.OrderResponse, error)
	CreateOrder(req *models.OrderRequest) (*models.OrderResponse, error)
	UpdateOrder(uuid string, req *models.OrderRequest) (*models.OrderResponse, error)
}

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &OrderRepo{db: config.DB}
}

func (or *OrderRepo) GetOrders() (*[]models.OrderResponse, error) {
	var orders []models.Order
	if err := or.db.Find(&orders).Error; err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return &[]models.OrderResponse{}, nil
	}

	var res []models.OrderResponse
	copier.Copy(&res, &orders)

	return &res, nil
}

func (pr *OrderRepo) GetOrderByID(uuid string) (*models.OrderResponse, error) {
	var order models.Order
	if err := pr.db.Where("uuid = ?", uuid).Find(&order).Error; err != nil {
		return nil, err
	}

	if order.UUID == "" {
		return nil, nil
	}

	var res models.OrderResponse
	copier.Copy(&res, &order)

	return &res, nil
}

func (pr *OrderRepo) CreateOrder(req *models.OrderRequest) (*models.OrderResponse, error) {
	var order models.Order

	copier.Copy(&order, &req)
	if err := pr.db.Create(&order).Error; err != nil {
		return nil, err
	}

	var res models.OrderResponse
	copier.Copy(&res, &order)

	return &res, nil
}

func (pr *OrderRepo) UpdateOrder(uuid string, req *models.OrderRequest) (*models.OrderResponse, error) {
	var order models.Order

	if err := pr.db.Where("uuid = ?", uuid).First(&order).Error; err != nil {
		return nil, err
	}

	order.Quantity = req.Quantity

	if err := pr.db.Save(&order).Error; err != nil {
		return nil, err
	}

	var res models.OrderResponse
	copier.Copy(&res, &order)

	return &res, nil
}

package repositories

import (
	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type CheckoutOrderRepository interface {
	GetCheckoutOrders() (*[]models.CheckoutOrderResponse, error)
	GetCheckoutOrderByID(uuid string) (*models.CheckoutOrderResponse, error)
	CreateCheckoutOrder(req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error)
	UpdateCheckoutOrder(uuid string, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error)
}

type CheckoutOrderRepo struct {
	db *gorm.DB
}

func NewCheckoutOrderRepository() CheckoutOrderRepository {
	return &CheckoutOrderRepo{db: config.DB}
}

func (cr *CheckoutOrderRepo) GetCheckoutOrders() (*[]models.CheckoutOrderResponse, error) {
	var checkoutOrders []models.CheckoutOrder
	if err := cr.db.Find(&checkoutOrders).Error; err != nil {
		return nil, err
	}

	if len(checkoutOrders) == 0 {
		return &[]models.CheckoutOrderResponse{}, nil
	}

	var res []models.CheckoutOrderResponse
	copier.Copy(&res, &checkoutOrders)

	return &res, nil
}

func (cr *CheckoutOrderRepo) GetCheckoutOrderByID(uuid string) (*models.CheckoutOrderResponse, error) {
	var checkoutOrders models.CheckoutOrder
	if err := cr.db.Where("uuid = ?", uuid).Find(&checkoutOrders).Error; err != nil {
		return nil, err
	}

	if checkoutOrders.UUID == "" {
		return nil, nil
	}

	var res models.CheckoutOrderResponse
	copier.Copy(&res, &checkoutOrders)

	return &res, nil
}
func (cr *CheckoutOrderRepo) CreateCheckoutOrder(req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error) {
	var checkoutOrder models.CheckoutOrder

	copier.Copy(&checkoutOrder, &req)
	if err := cr.db.Create(&checkoutOrder).Error; err != nil {
		return nil, err
	}

	var orders models.CheckoutOrder
	var totalPrice uint
	for _, order := range orders.Orders {
		order.UUID = checkoutOrder.UUID
		if err := cr.db.Create(&order).Error; err != nil {
			return nil, err
		}
		orders.Orders = append(orders.Orders, order)
		totalPrice += order.Price
	}
	checkoutOrder.TotalPrice = totalPrice
	checkoutOrder.Orders = orders.Orders

	if err := cr.db.Save(&checkoutOrder).Error; err != nil {
		return nil, err
	}

	var res models.CheckoutOrderResponse
	copier.Copy(&res, &checkoutOrder)

	return &res, nil
}
func (cr *CheckoutOrderRepo) UpdateCheckoutOrder(uuid string, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error) {
	var checkoutOrders models.CheckoutOrder

	if err := cr.db.Where("uuid = ?", uuid).First(&checkoutOrders).Error; err != nil {
		return nil, err
	}

	checkoutOrders.Orders = req.Orders

	if err := cr.db.Save(&checkoutOrders).Error; err != nil {
		return nil, err
	}

	var res models.CheckoutOrderResponse
	copier.Copy(&res, &checkoutOrders)

	return &res, nil
}

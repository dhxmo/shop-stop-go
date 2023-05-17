package services

import (
	"context"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type OrderService interface {
	GetOrders(ctx context.Context) (*[]models.OrderResponse, error)
	GetOrderByID(ctx context.Context, orderID string) (*models.OrderResponse, error)
	CreateOrder(ctx context.Context, req *models.OrderRequest) (*models.OrderResponse, error)
	UpdateOrder(ctx context.Context, uuid string, req *models.OrderRequest) (*models.OrderResponse, error)
}

type OrderSvc struct {
	repo repositories.OrderRepository
}

func NewOrderSvc() OrderService {
	return &OrderSvc{repo: repositories.NewOrderRepository()}
}

func (os *OrderSvc) GetOrders(ctx context.Context) (*[]models.OrderResponse, error) {
	orders, err := os.repo.GetOrders()
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return orders, nil
}

func (os *OrderSvc) GetOrderByID(ctx context.Context, orderID string) (*models.OrderResponse, error) {
	order, err := os.repo.GetOrderByID(orderID)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return order, nil
}

func (os *OrderSvc) CreateOrder(ctx context.Context, req *models.OrderRequest) (*models.OrderResponse, error) {
	order, err := os.repo.CreateOrder(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return order, nil
}

func (os *OrderSvc) UpdateOrder(ctx context.Context, uuid string, req *models.OrderRequest) (*models.OrderResponse, error) {
	order, err := os.repo.UpdateOrder(uuid, req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return order, nil

}

package services

import (
	"context"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type CheckoutService interface {
	GetCheckouts(ctx context.Context) (*[]models.CheckoutOrderResponse, error)
	GetCheckoutByID(ctx context.Context, checkoutOrdersID string) (*models.CheckoutOrderResponse, error)
	CreateCheckout(ctx context.Context, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error)
	UpdateCheckout(ctx context.Context, uuid string, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error)
}

type CheckoutSvc struct {
	repo repositories.CheckoutOrderRepository
}

func NewCheckoutSvc() CheckoutService {
	return &CheckoutSvc{repo: repositories.NewCheckoutOrderRepository()}
}

func (cs *CheckoutSvc) GetCheckouts(ctx context.Context) (*[]models.CheckoutOrderResponse, error) {
	checkoutOrders, err := cs.repo.GetCheckoutOrders()
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return checkoutOrders, nil
}

func (cs *CheckoutSvc) GetCheckoutByID(ctx context.Context, checkoutOrdersID string) (*models.CheckoutOrderResponse, error) {
	checkoutOrders, err := cs.repo.GetCheckoutOrderByID(checkoutOrdersID)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return checkoutOrders, nil
}

func (cs *CheckoutSvc) CreateCheckout(ctx context.Context, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error) {
	checkout, err := cs.repo.CreateCheckoutOrder(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return checkout, nil
}

func (cs *CheckoutSvc) UpdateCheckout(ctx context.Context, uuid string, req *models.CheckoutOrderBodyRequest) (*models.CheckoutOrderResponse, error) {
	checkout, err := cs.repo.UpdateCheckoutOrder(uuid, req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return checkout, nil

}

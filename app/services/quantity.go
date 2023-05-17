package services

import (
	"context"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/app/repositories"
)

type QuantityService interface {
	GetQuantities(ctx context.Context) (*[]models.QuantityResponse, error)
	GetQuantityByID(ctx context.Context, quantitiesID string) (*models.QuantityResponse, error)
	CreateQuantity(ctx context.Context, req *models.QuantityBodyRequest) (*models.QuantityResponse, error)
	UpdateQuantity(ctx context.Context, uuid string, req *models.QuantityBodyRequest) (*models.QuantityResponse, error)
}

type QuantitySvc struct {
	repo repositories.QuantityRepository
}

func NewQuantitySvc() QuantityService {
	return &QuantitySvc{repo: repositories.NewQuantityRepository()}
}

func (qs *QuantitySvc) GetQuantities(ctx context.Context) (*[]models.QuantityResponse, error) {
	quantities, err := qs.repo.GetQuantities()
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return quantities, nil
}

func (qs *QuantitySvc) GetQuantityByID(ctx context.Context, quantitiesID string) (*models.QuantityResponse, error) {
	quantity, err := qs.repo.GetQuantityByID(quantitiesID)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return quantity, nil

}

func (qs *QuantitySvc) CreateQuantity(ctx context.Context, req *models.QuantityBodyRequest) (*models.QuantityResponse, error) {
	quantity, err := qs.repo.CreateQuantity(req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return quantity, nil
}

func (qs *QuantitySvc) UpdateQuantity(ctx context.Context, uuid string, req *models.QuantityBodyRequest) (*models.QuantityResponse, error) {
	quantity, err := qs.repo.UpdateQuantity(uuid, req)
	if err != nil {
		ctx.Err()
		return nil, err
	}
	return quantity, nil
}

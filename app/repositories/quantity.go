package repositories

import (
	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type QuantityRepository interface {
	GetQuantities() (*[]models.QuantityResponse, error)
	GetQuantityByID(uuid string) (*models.QuantityResponse, error)
	CreateQuantity(req *models.QuantityBodyRequest) (*models.QuantityResponse, error)
	UpdateQuantity(uuid string, req *models.QuantityBodyRequest) (*models.QuantityResponse, error)
}

type QuantityRepo struct {
	db *gorm.DB
}

func NewQuantityRepository() QuantityRepository {
	return &QuantityRepo{db: config.DB}
}

func (cr *QuantityRepo) GetQuantities() (*[]models.QuantityResponse, error) {
	var quantities []models.Quantity
	if err := cr.db.Find(&quantities).Error; err != nil {
		return nil, err
	}

	if len(quantities) == 0 {
		return &[]models.QuantityResponse{}, nil
	}

	var res []models.QuantityResponse
	copier.Copy(&res, &quantities)

	return &res, nil
}

func (cr *QuantityRepo) GetQuantityByID(uuid string) (*models.QuantityResponse, error) {
	var quantity models.Quantity
	if err := cr.db.Where("uuid = ?", uuid).Find(&quantity).Error; err != nil {
		return nil, err
	}

	if quantity.UUID == "" {
		return nil, nil
	}

	var res models.QuantityResponse
	copier.Copy(&res, &quantity)

	return &res, nil
}

func (cr *QuantityRepo) CreateQuantity(req *models.QuantityBodyRequest) (*models.QuantityResponse, error) {
	var quantity models.Quantity

	copier.Copy(&quantity, &req)
	if err := cr.db.Create(&quantity).Error; err != nil {
		return nil, err
	}

	var res models.QuantityResponse
	copier.Copy(&res, &quantity)

	return &res, nil
}

func (cr *QuantityRepo) UpdateQuantity(uuid string, req *models.QuantityBodyRequest) (*models.QuantityResponse, error) {
	var quantities models.Quantity

	if err := cr.db.Where("uuid = ?", uuid).First(&quantities).Error; err != nil {
		return nil, err
	}

	quantities.Quantity = req.Quantity

	if err := cr.db.Save(&quantities).Error; err != nil {
		return nil, err
	}

	var res models.QuantityResponse
	copier.Copy(&res, &quantities)

	return &res, nil
}

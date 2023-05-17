package services

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategorySvc)
	_ = container.Provide(NewOrderSvc)
	_ = container.Provide(NewProductSvc)
	_ = container.Provide(NewQuantitySvc)
	_ = container.Provide(NewUserSvc)
	_ = container.Provide(NewCheckoutSvc)
	return nil
}

package controllers

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryController)
	_ = container.Provide(NewProductController)
	_ = container.Provide(NewQuantityController)
	_ = container.Provide(NewCheckoutController)
	_ = container.Provide(NewOrderController)
	_ = container.Provide(NewUserController)
	return nil
}

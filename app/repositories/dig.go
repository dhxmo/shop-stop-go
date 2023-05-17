package repositories

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryRepository)
	_ = container.Provide(NewProductRepository)
	_ = container.Provide(NewOrderRepository)
	_ = container.Provide(NewCheckoutOrderRepository)
	_ = container.Provide(NewQuantityRepository)
	_ = container.Provide(NewUserRepository)
	return nil
}

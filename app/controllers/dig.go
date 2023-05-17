package controllers

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryController)
	// _ = container.Provide(NewProductAPI)
	return nil
}

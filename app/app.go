package app

import (
	"log"

	"go.uber.org/dig"

	"github.com/dhxmo/shop-stop-go/app/controllers"
	"github.com/dhxmo/shop-stop-go/app/repositories"
	"github.com/dhxmo/shop-stop-go/app/services"
)

func Init() {
	BuildContainer()
}

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := repositories.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject repositories", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject services", err)
	}

	// Inject APIs
	err = controllers.Inject(container)
	if err != nil {
		log.Fatal("Failed to inject controllers", err)
	}

	return container
}

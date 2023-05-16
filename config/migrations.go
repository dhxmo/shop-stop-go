package config

import (
	"log"

	"github.com/dhxmo/shop-stop-go/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func Migrate() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("error in db connection", err)
	}
	log.Println("successful db connection")

	Product := models.Product{}
	Category := models.Category{}
	CheckoutOrder := models.CheckoutOrder{}
	Order := models.Order{}
	User := models.User{}

	db.AutoMigrate(&Product, &Category, &CheckoutOrder, &Order, &User)
}
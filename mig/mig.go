package mig

import (
	"myproject/configs"
	"myproject/models"
)

func Migrate() {

	var DBCon = configs.CreateCon()
	DBCon.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Category{},
		&models.Product{},
		&models.Attribute{},
		&models.Vendor{},
		&models.AttributeProduct{},
		&models.ProductVendor{},
		&models.Comment{},
		&models.Order{},
		&models.Transaction{},
		&models.OrderStatus{},
		&models.Shipper{},
		&models.Zipcode{},
		&models.Cart{},
	)
}

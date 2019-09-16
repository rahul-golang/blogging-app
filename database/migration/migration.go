package main

import (
	"blogging-app/pkg/models"
	"log"

	"github.com/go-gormigrate/gormigrate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// ID: "201608301400",
			// Migrate: func(tx *gorm.DB) error {
			// 	return tx.AutoMigrate(&products_models.Product{},
			// 		&products_models.ProductDescription{},
			// 		&products_models.ProductImages{},
			// 	).Error
			// },
			// Rollback: func(tx *gorm.DB) error {
			// 	return tx.DropTable("orders", "product_descriptions", "product_images").Error
			// },
		},
	})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
		).Error
		if err != nil {
			return err
		}

		// if err := tx.Model(products_models.ProductImages{}).AddForeignKey("products_id", "products(id)", "CASCADE", "CASCADE").Error; err != nil {
		// 	return err
		// }
		// if err := tx.Model(products_models.ProductDescription{}).AddForeignKey("products_id", "products(id)", "CASCADE", "CASCADE").Error; err != nil {
		// 	return err
		// }

		// if err := tx.Model(orders_models.OrderProducts{}).AddForeignKey("orders_id", "orders (orders_id)", "CASCADE", "CASCADE").Error; err != nil {
		// 	return err
		// }
		// if err := tx.Model(orders_models.Orders{}).AddForeignKey("order_status", "order_statuses(order_status_id)", "CASCADE", "CASCADE").Error; err != nil {
		// 	return err
		// }

		// all other foreign keys...
		return nil
	})

	err = m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}
}

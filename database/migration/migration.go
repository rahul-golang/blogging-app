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
			&models.Blog{},
			&models.Followers{},
		).Error
		if err != nil {
			return err
		}

		if err := tx.Model(models.Blog{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
			return err
		}
		if err := tx.Model(models.Followers{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
			return err
		}
		if err := tx.Model(models.Followers{}).AddForeignKey("follower_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
			return err
		}

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

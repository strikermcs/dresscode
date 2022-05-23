package config

import (
	"log"

	"github.com/strikermcs/dresscode/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	
	db, err := gorm.Open(sqlite.Open("dresscode.db"), &gorm.Config{})

	if err != nil {
    	panic("failed to connect database")
	}
	
	log.Println("Connected to Database Successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")

	db.AutoMigrate(&models.Product{}, &models.ProductDescription{}, &models.ProductImage{}, &models.ProductSize{})
	
	Database = DbInstance{
		Db: db,
	}

} 
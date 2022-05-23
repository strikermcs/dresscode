package controllers

import (
	"github.com/strikermcs/dresscode/pkg/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase(base *gorm.DB){
	db = base

	db.AutoMigrate(
		&models.Product{}, 
		&models.ProductDescription{})
} 
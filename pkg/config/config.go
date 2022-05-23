package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func GetDb() *gorm.DB {
	
	db, err := gorm.Open(sqlite.Open("dresscode.db"), &gorm.Config{})

	if err != nil {
    	panic("failed to connect database")
	}

	return db
} 
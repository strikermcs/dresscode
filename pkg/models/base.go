package models

import (
	"time"
	"github.com/satori/go.uuid"	
	"gorm.io/gorm"
)

type Base struct {
	ID 			uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt 	time.Time `json:"created_at"` 	
	UpdatedAt 	time.Time `json:"updated_at"`	
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()
	return
}

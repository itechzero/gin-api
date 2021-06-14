package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"column:id;primaryKey" json:"-" form:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

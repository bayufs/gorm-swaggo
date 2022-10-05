package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Order struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key,column:id"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	IsDel        soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New()
	order.CreatedAt = time.Now()
	return nil
}

func (order *Order) BeforeUpdate(tx *gorm.DB) error {
	order.UpdatedAt = time.Now()
	return nil
}

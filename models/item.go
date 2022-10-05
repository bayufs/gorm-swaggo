package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key,column:id,constraint:OnDelete:CASCADE"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     uuid.UUID
}

func (item *Item) BeforeCreate(tx *gorm.DB) error {
	item.ID = uuid.New()
	return nil
}

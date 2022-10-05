package repository

import (
	"fmt"
	"task-sesi8/configs"
	"task-sesi8/dto"
	"task-sesi8/models"
	"time"
)

type OrderRepository interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrders() ([]dto.Order, error)
	GetOrderById(id string) (models.Order, error)
	UpdateOrder(id string, order models.Order) (dto.Order, error)
	DeleteOrder(id string) (any, error)
}

func NewOrderRepository() OrderRepository {
	return &dbConnection{
		connection: configs.ConnectDB(),
	}
}

func (db *dbConnection) CreateOrder(order models.Order) (models.Order, error) {
	fmt.Println(order)
	err := db.connection.Model(&models.Order{}).Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (db *dbConnection) GetOrders() ([]dto.Order, error) {
	var orders []dto.Order
	err := db.connection.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (db *dbConnection) GetOrderById(id string) (models.Order, error) {
	var order models.Order
	err := db.connection.Preload("Items").Find(&order, id).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (db *dbConnection) UpdateOrder(id string, data models.Order) (dto.Order, error) {
	var order dto.Order
	updateData := make(map[string]interface{})

	updateData["customer_name"] = data.CustomerName

	err := db.connection.Model(&models.Order{}).Where("id = ?", id).Updates(&updateData).Error
	if err != nil {
		return order, err
	}

	err = db.connection.Model(&models.Order{}).Where("id = ?", id).Preload("Items").Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (db *dbConnection) DeleteOrder(id string) (any, error) {
	fmt.Println("ruun", id)
	deleteData := make(map[string]interface{})

	deleteData["is_del"] = 1
	deleteData["deleted_at"] = time.Now()
	res := db.connection.Model(&models.Order{}).Where("id = ?", id).Updates(&deleteData)
	if res.Error != nil {
		return nil, res.Error
	}
	return res.RowsAffected, nil
}

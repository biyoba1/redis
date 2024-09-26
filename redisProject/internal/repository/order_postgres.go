package repository

import (
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/internal/models"
	"gorm.io/gorm"
	"log"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewOrdersPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateOrder(order models.Order) (int, error) {
	id := initializer.DB.Create(&order)
	if id.Error != nil {
		return 0, id.Error
	}
	//redis.Redis(order.Products)
	return int(id.RowsAffected), nil
}

func (r *TodoItemPostgres) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := initializer.DB.Preload("Person").Preload("Products").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *TodoItemPostgres) GetOrderByID(id int) (models.Order, error) {
	order := models.Order{}

	err := initializer.DB.Preload("Person").Preload("Products").First(&order, id).Error
	if err != nil {
		log.Println(err)
		return models.Order{}, err
	}

	return order, nil
}

func (r *TodoItemPostgres) DeleteOrder(orderId int) error {
	var order models.Order
	err := initializer.DB.Delete(&order, orderId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoItemPostgres) UpdateOrder(orderId int, order models.UpdateOrderInput) error {
	return nil
}

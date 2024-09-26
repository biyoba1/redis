package services

import (
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/biyoba1/redisProject/internal/repository"
)

type OrderStruct struct {
	repo repository.OrderService
}

func NewOrderService(repo repository.OrderService) *OrderStruct {
	return &OrderStruct{repo: repo}
}

func (o *OrderStruct) CreateOrder(order models.Order) (int, error) {
	return o.repo.CreateOrder(order)
}

func (o *OrderStruct) GetAllOrders() ([]models.Order, error) {
	return o.repo.GetAllOrders()
}

func (o *OrderStruct) GetOrderByID(id int) (models.Order, error) {
	return o.repo.GetOrderByID(id)
}

func (o *OrderStruct) DeleteOrder(orderId int) error {
	return o.repo.DeleteOrder(orderId)
}

func (o *OrderStruct) UpdateOrder(orderId int, order models.UpdateOrderInput) error {
	return o.repo.UpdateOrder(orderId, order)
}

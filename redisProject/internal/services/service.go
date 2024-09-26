package services

import (
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/biyoba1/redisProject/internal/repository"
)

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id int) (models.Order, error)
	CreateOrder(order models.Order) (int, error)
	UpdateOrder(orderId int, order models.UpdateOrderInput) error
	DeleteOrder(orderId int) error
}

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetByNameProduct(name string) (models.Product, error)
	CreateProduct(product models.Product) (int, error)
	UpdateProduct(productId int, product models.Product) error
	DeleteProduct(productId int) error
}

type Service struct {
	OrderService
	ProductService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		OrderService:   NewOrderService(repos.OrderService),
		ProductService: NewProductService(repos.ProductService),
	}
}

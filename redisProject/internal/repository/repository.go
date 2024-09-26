package repository

import (
	"github.com/biyoba1/redisProject/internal/models"
	"gorm.io/gorm"
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

type Repository struct {
	OrderService
	ProductService
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		OrderService:   NewOrdersPostgres(db),
		ProductService: NewProductsPostgres(db),
	}
}

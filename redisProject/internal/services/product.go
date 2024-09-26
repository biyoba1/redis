package services

import (
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/biyoba1/redisProject/internal/repository"
)

type ProductStruct struct {
	repo repository.ProductService
}

func NewProductService(repo repository.ProductService) *ProductStruct {
	return &ProductStruct{repo: repo}
}

func (s *ProductStruct) CreateProduct(product models.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductStruct) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductStruct) GetByNameProduct(name string) (models.Product, error) {
	return s.repo.GetByNameProduct(name)
}

func (s *ProductStruct) DeleteProduct(productId int) error {
	return s.repo.DeleteProduct(productId)
}

func (s *ProductStruct) UpdateProduct(productId int, product models.Product) error {
	return s.repo.UpdateProduct(productId, product)
}

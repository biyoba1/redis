package repository

import (
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/internal/models"
	"gorm.io/gorm"
	"log"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductsPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(product models.Product) (int, error) {
	id := initializer.DB.Create(&product)
	if id.Error != nil {
		return 0, id.Error
	}
	return int(id.RowsAffected), nil
}

func (r *ProductPostgres) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := initializer.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductPostgres) GetByNameProduct(name string) (models.Product, error) {
	product := models.Product{}
	err := initializer.DB.First(&product, "name=?", name).Error
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	log.Println("Продукт получен из Базы данных")

	product.Counter++
	err = initializer.DB.Save(&product).Error
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductPostgres) DeleteProduct(productId int) error {
	var product models.Product
	err := initializer.DB.Delete(&product, productId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductPostgres) UpdateProduct(productId int, product models.Product) error {
	return nil
}

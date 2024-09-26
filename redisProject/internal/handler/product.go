package handler

import (
	"fmt"
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/biyoba1/redisProject/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var input models.Product
	err := c.Bind(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	id, err := h.services.ProductService.CreateProduct(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (h *Handler) GetByNameProduct(c *gin.Context) {
	name := c.Query("name")

	product, err := redis.GetCacheProduct(name)

	if err == nil {
		fmt.Println("Продукт получен из кэша")
	}

	if err != nil {
		if err.Error() == fmt.Sprintf("Key %s not found in redis cache", name) {
			product, err = h.services.ProductService.GetByNameProduct(name)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return // Добавьте ключевое слово return здесь
			}
			err = redis.CacheProduct(product)
			if err != nil {
				log.Println(err)
			}
		} else {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return // Добавьте ключевое слово return здесь
		}
	} else if product.Name == "" {
		newErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.services.ProductService.GetAllProducts()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Products not found")
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) UpdateProduct(c *gin.Context) {

}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
	}
	err = h.services.ProductService.DeleteProduct(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

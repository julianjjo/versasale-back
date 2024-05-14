package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	productId := c.Param("productId")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":            productId,
			"name":          "Producto 1",
			"price":         1000,
			"stock":         10,
			"description":   "Producto 1 descripción",
			"product_image": "https://via.placeholder.com/200",
			"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
		},
	},
	)
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": []gin.H{
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Producto 1",
				"price":         1000,
				"stock":         10,
				"description":   "Producto 1 descripción",
				"product_image": "https://via.placeholder.com/200",
				"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 2",
				"price":         2000,
				"stock":         20,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 2 description",
				"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 3",
				"price":         3000,
				"stock":         30,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 3 description",
				"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 4",
				"price":         4000,
				"stock":         40,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 4 description",
				"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
			},
			{
				"id":            "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
				"name":          "Product 5",
				"price":         5000,
				"stock":         50,
				"product_image": "https://via.placeholder.com/200",
				"description":   "Product 5 description",
				"category":      "ce2d247f-80a8-4ca5-a646-90ec49311a7f",
			},
		},
	},
	)
}

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}

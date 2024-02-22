package main

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.Default()

	customers := router.Group("/customers")
	{
		customers.GET("", c.CustomerController.ReadAllCustomer)
		customers.GET("/:customerID", c.CustomerController.ReadCustomerByID)
		customers.POST("", c.CustomerController.CreateCustomer)
		customers.PUT("/:customerID", c.CustomerController.UpdateCustomer)
		customers.DELETE("/:customerID", c.CustomerController.DeleteCustomer)
	}

	categories := router.Group("/categories")
	{
		categories.GET("", c.CategoryController.ReadAllCategory)
		categories.GET("/:categoryID", c.CategoryController.ReadCategoryByID)
		categories.POST("", c.CategoryController.CreateCategory)
		categories.PUT("", c.CategoryController.UpdateCategory)
		categories.DELETE("/:categoryID", c.CategoryController.DeleteCategory)
	}

	products := router.Group("/products")
	{
		products.GET("", c.ProductController.GetProductsWithSupplier)
		products.GET("/above-average-price", c.ProductController.GetProductAboveAveragePrice)
	}

	return router
}

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
	}

	return router
}

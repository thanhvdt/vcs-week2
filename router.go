package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhvdt/vcs-week2/controller"
)

func NewRouter(c *controller.CustomerController) *gin.Engine {
	router := gin.Default()
	router.GET("/customers", c.ReadAllCustomer)
	router.GET("/customers/:customerID", c.ReadCustomerByID)
	router.POST("/customers", c.CreateCustomer)
	router.PUT("/customers/:customerID", c.UpdateCustomer)
	router.DELETE("/customers/:customerID", c.DeleteCustomer)
	return router
}

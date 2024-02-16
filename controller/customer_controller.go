package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	customer2 "github.com/thanhvdt/vcs-week2/data/request/customer"
	"github.com/thanhvdt/vcs-week2/data/response"
	"github.com/thanhvdt/vcs-week2/service/customer"
	"net/http"
)

type CustomerController struct {
	CustomerService customer.CustomerService
}

func NewCustomerController(customerService customer.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	log.Info().Msg("Creating Customer")
	var createCustomerRequest = customer2.CreateCustomerRequest{}
	err := ctx.ShouldBindJSON(&createCustomerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	customer, er := c.CustomerService.Create(createCustomerRequest)
	if er != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Created",
		Data:   customer,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, serverResponse)

}

func (c *CustomerController) ReadAllCustomer(ctx *gin.Context) {
	log.Info().Msg("Reading All Customer")
	customers, err := c.CustomerService.ReadAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customers,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, serverResponse)
}

func (c *CustomerController) ReadCustomerByID(ctx *gin.Context) {
	log.Info().Msg("Reading Customer By ID")
	customerID := ctx.Param("customerID")
	customer, err := c.CustomerService.ReadByID(customerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customer,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, serverResponse)
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	log.Info().Msg("Updating Customer")
	customerID := ctx.Param("customerID")
	var updateCustomerRequest = customer2.UpdateCustomerRequest{}
	err := ctx.ShouldBindJSON(&updateCustomerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	customer, er := c.CustomerService.Update(customerID, updateCustomerRequest)
	if er != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   customer,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, serverResponse)
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	log.Info().Msg("Deleting Customer")
	customerID := ctx.Param("customerID")
	err := c.CustomerService.Delete(customerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	serverResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Deleted",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, serverResponse)
}

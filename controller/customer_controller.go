package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	customer2 "github.com/thanhvdt/vcs-week2/data/request/customer"
	"github.com/thanhvdt/vcs-week2/data/response"
	_ "github.com/thanhvdt/vcs-week2/docs"
	"github.com/thanhvdt/vcs-week2/service/customer"
	"net/http"
)

type CustomerController struct {
	CustomerService customer.CustomerService
}

func NewCustomerController(customerService customer.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

// CreateCustomer
// @Tags Customer
// @Summary Create a new customer
// @Description Create a new customer
// @Accept json
// @Produce json
// @Param customer body customer2.CreateCustomerRequest true "Customer object that needs to be added"
// @Success 200 {object} response.Response
// @Router /customers [post]
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

// ReadAllCustomer
// @Tags Customer
// @Summary Read all customers
// @Description Read all customers
// @Produce json
// @Success 200 {object} response.Response
// @Router /customers [get]
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

// ReadCustomerByID
// @Tags Customer
// @Summary Read customer by ID
// @Description Read customer by ID
// @Produce json
// @Param customerID path string true "Customer ID"
// @Success 200 {object} response.Response
// @Router /customers/{customerID} [get]
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

// UpdateCustomer
// @Tags Customer
// @Summary Update customer
// @Description Update customer
// @Accept json
// @Produce json
// @Param customerID path string true "Customer ID"
// @Param customer body customer2.UpdateCustomerRequest true "Customer object that needs to be updated"
// @Success 200 {object} response.Response
// @Router /customers/{customerID} [put]
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

// DeleteCustomer
// @Tags Customer
// @Summary Delete customer
// @Description Delete customer
// @Produce json
// @Param customerID path string true "Customer ID"
// @Success 200 {object} response.Response
// @Router /customers/{customerID} [delete]
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

// SearchCustomerByCompany
// @Tags Customer
// @Summary Search customer by company
// @Description Search customer by company
// @Produce json
// @Param company path string true "Company"
// @Success 200 {object} response.Response
// @Router /customers/search-by-company/{company} [get]
func (c *CustomerController) SearchCustomerByCompany(ctx *gin.Context) {
	log.Info().Msg("Searching Customer By Company")
	company := ctx.Param("company")
	customers, err := c.CustomerService.SearchByCompany(company)
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

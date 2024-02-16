package customer

import (
	"github.com/go-playground/validator/v10"
	customer2 "github.com/thanhvdt/vcs-week2/data/request/customer"
	"github.com/thanhvdt/vcs-week2/data/response"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository/customer"
)

type CustomerServiceImpl struct {
	CustomerRepository customer.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository customer.CustomerRepository, validate *validator.Validate) *CustomerServiceImpl {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, Validate: validate}
}

func (c *CustomerServiceImpl) Create(customer customer2.CreateCustomerRequest) (*model.Customer, error) {
	err := c.Validate.Struct(customer)
	if err != nil {
		return nil, err
	}
	customerModel := model.Customer{
		CustomerID:   customer.CustomerID,
		CompanyName:  customer.CompanyName,
		ContactName:  customer.ContactName,
		ContactTitle: customer.ContactTitle,
		Address:      customer.Address,
		City:         customer.City,
		Region:       customer.Region,
		PostalCode:   customer.PostalCode,
		Country:      customer.Country,
		Phone:        customer.Phone,
		Fax:          customer.Fax,
	}
	return c.CustomerRepository.Save(&customerModel)
}

func (c *CustomerServiceImpl) ReadAll() ([]response.CustomerResponse, error) {
	var customers []response.CustomerResponse
	allCustomer, err := c.CustomerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	for _, customer := range allCustomer {
		customers = append(customers, response.CustomerResponse{
			CustomerID:   customer.CustomerID,
			CompanyName:  customer.CompanyName,
			ContactName:  customer.ContactName,
			ContactTitle: customer.ContactTitle,
			Address:      customer.Address,
			City:         customer.City,
			Region:       customer.Region,
			PostalCode:   customer.PostalCode,
			Country:      customer.Country,
			Phone:        customer.Phone,
			Fax:          customer.Fax,
		})
	}

	return customers, nil
}

func (c *CustomerServiceImpl) ReadByID(customerID string) (response.CustomerResponse, error) {
	customer, err := c.CustomerRepository.FindByID(customerID)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	return response.CustomerResponse{
		CustomerID:   customer.CustomerID,
		CompanyName:  customer.CompanyName,
		ContactName:  customer.ContactName,
		ContactTitle: customer.ContactTitle,
		Address:      customer.Address,
		City:         customer.City,
		Region:       customer.Region,
		PostalCode:   customer.PostalCode,
		Country:      customer.Country,
		Phone:        customer.Phone,
		Fax:          customer.Fax,
	}, nil
}

func (c *CustomerServiceImpl) Update(customerID string, customer customer2.UpdateCustomerRequest) (*model.Customer, error) {
	err := c.Validate.Struct(customer)
	if err != nil {
		return nil, err
	}
	customerModel := model.Customer{
		CustomerID:   customerID,
		CompanyName:  customer.CompanyName,
		ContactName:  customer.ContactName,
		ContactTitle: customer.ContactTitle,
		Address:      customer.Address,
		City:         customer.City,
		Region:       customer.Region,
		PostalCode:   customer.PostalCode,
		Country:      customer.Country,
		Phone:        customer.Phone,
		Fax:          customer.Fax,
	}
	return c.CustomerRepository.Update(&customerModel)
}

func (c *CustomerServiceImpl) Delete(customerID string) error {
	return c.CustomerRepository.Delete(customerID)
}

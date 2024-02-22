package customer

import (
	"github.com/thanhvdt/vcs-week2/data/request/customer"
	"github.com/thanhvdt/vcs-week2/data/response"
	"github.com/thanhvdt/vcs-week2/model"
)

type CustomerService interface {
	Create(customer customer.CreateCustomerRequest) (*model.Customer, error)
	ReadAll() ([]response.CustomerResponse, error)
	ReadByID(customerID string) (response.CustomerResponse, error)
	Update(customerID string, customer customer.UpdateCustomerRequest) (*model.Customer, error)
	Delete(customerID string) error
	SearchByCompany(company string) (map[string]interface{}, error)
}

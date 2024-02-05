package service

import (
	"github.com/thanhvdt/vcs-week2/data/request"
	"github.com/thanhvdt/vcs-week2/data/response"
	"github.com/thanhvdt/vcs-week2/model"
)

type CustomerService interface {
	Create(customer request.CreateCustomerRequest) (*model.Customer, error)
	ReadAll() ([]response.CustomerResponse, error)
	ReadByID(customerID string) (response.CustomerResponse, error)
	Update(customerID string, customer request.UpdateCustomerRequest) (*model.Customer, error)
	Delete(customerID string) error
}
